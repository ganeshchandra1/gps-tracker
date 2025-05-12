<script setup>
import { ref, onMounted, onBeforeUnmount, watch } from 'vue'
import mapboxgl from 'mapbox-gl'
import 'mapbox-gl/dist/mapbox-gl.css'
const API_BASE_URL = import.meta.env.VITE_BACKEND_URL
const map = ref(null)
const mapContainer = ref(null)
const markers = ref([])
const routeLayers = ref([])

const vehicles = ref([])
const filteredVehicles = ref([])

const selectedStatus = ref('All')
const selectedType = ref('All')
const showOnlineOnly = ref(false)
const showVehiclesPanel = ref(false)
const selectedVehicleId = ref(null)

let intervalId = null

const applyFilters = () => {
  filteredVehicles.value = vehicles.value.filter(v =>
    (selectedStatus.value === 'All' || v.status === selectedStatus.value) &&
    (selectedType.value === 'All' || v.vehicleType === selectedType.value) &&
    (!showOnlineOnly.value || v.isOnline)
  )
}

const fetchVehiclesAndPlot = async () => {
  try {
    await fetch(`${API_BASE_URL}vehicles/random-update`, { method: 'POST' })

    const response = await fetch(`${API_BASE_URL}vehicles`)
    const data = await response.json()
    vehicles.value = data
    applyFilters()

    markers.value.forEach(marker => marker.remove())
    markers.value = []
    routeLayers.value.forEach(id => {
      if (map.value.getLayer(id)) map.value.removeLayer(id)
      if (map.value.getSource(id)) map.value.removeSource(id)
    })
    routeLayers.value = []

    filteredVehicles.value.forEach(vehicle => {
      const marker = new mapboxgl.Marker({
        color: vehicle.status === 'Moving' ? 'green' :
               vehicle.status === 'Idle' ? 'yellow' : 'red'
      })
        .setLngLat([vehicle.longitude, vehicle.latitude])
        .setPopup(new mapboxgl.Popup().setHTML(
            `<div class="w-full p-2 bg-white rounded shadow-md text-sm text-black">
      <div class="font-bold text-lg mb-1">${vehicle.id}</div>
      <div>Driver: <span class="font-semibold">${vehicle.driverName}</span></div>
      <div>Status: 
        <span class="inline-block px-2 py-0.5 text-xs rounded-full 
          ${vehicle.status === 'Moving' ? 'bg-green-200 text-green-800' : 
            vehicle.status === 'Idle' ? 'bg-yellow-200 text-yellow-800' : 
            'bg-red-200 text-red-800'}">
          ${vehicle.status}
        </span>
      </div>
    </div>`
        ))
        .addTo(map.value)
      markers.value.push(marker)

      marker.getElement().addEventListener('click', () => {
        map.value.flyTo({
          center: [vehicle.longitude, vehicle.latitude],
          zoom: 16,
          speed: 1.2,
          curve: 1.4
        })
      })

      if (vehicle.routeCoordinates && vehicle.routeCoordinates.length > 1) {
        const coordinates = vehicle.routeCoordinates.map(coord => [coord.longitude, coord.latitude])
        const routeId = `route-${vehicle.id}`

        map.value.addSource(routeId, {
          type: 'geojson',
          data: {
            type: 'Feature',
            geometry: { type: 'LineString', coordinates }
          }
        })

        map.value.addLayer({
          id: routeId,
          type: 'line',
          source: routeId,
          layout: { 'line-join': 'round', 'line-cap': 'round' },
          paint: { 'line-color': '#22c55e', 'line-width': 4 }
        })

        routeLayers.value.push(routeId)
      }
    })
  } catch (error) {
    console.error('Failed to load vehicles:', error)
  }
}

const expandVehicle = (vehicle) => {
  selectedVehicleId.value = selectedVehicleId.value === vehicle.id ? null : vehicle.id
  map.value.flyTo({
    center: [vehicle.longitude, vehicle.latitude],
    zoom: 16,
    speed: 1.2,
    curve: 1.4
  })
  markers.value.forEach(marker => marker.getElement().style.transform = 'scale(1)')
  const marker = markers.value.find(m => {
    const lngLat = m.getLngLat()
    return lngLat.lng === vehicle.longitude && lngLat.lat === vehicle.latitude
  })
  if (marker) marker.getElement().style.transform = 'scale(1.5)'
}
watch([selectedStatus, selectedType, showOnlineOnly], applyFilters)
onMounted(() => {
  mapboxgl.accessToken = 'pk.eyJ1IjoiZ2FuaTIwOCIsImEiOiJjbWFrN2lkYWgwNTByMnNwdzI4eTF6cjFiIn0.AbTOPiSgdJaLIyjJpw4LcA'
  map.value = new mapboxgl.Map({
    container: mapContainer.value,
    style: 'mapbox://styles/mapbox/streets-v11',
    center: [-98.35, 39.5],
    zoom: 4
  })
  map.value.resize()
  map.value.addControl(new mapboxgl.NavigationControl(), 'top-right')
  fetchVehiclesAndPlot()
})
onBeforeUnmount(() => {
  clearInterval(intervalId)
  map.value?.remove()
})
const resetFilters = () => {
  selectedStatus.value = 'All'
  selectedType.value = 'All'
  showOnlineOnly.value = false
}
</script>
<template>
<div class="w-full h-screen">
  <div class="bg-gray-600 p-3 shadow-md flex flex-wrap items-center justify-start gap-4">
    <div class="bg-gray-800 p-2">
      <label class="text-sm text-white font-semibold">Status:</label>
      <select v-model="selectedStatus" class="border-none outline-none focus:ring-0 focus:outline-none bg-gray-800 text-white rounded px-2 py-1">
        <option>All</option>
        <option>Moving</option>
        <option>Idle</option>
        <option>Stopped</option>
      </select>
    </div>

    <div class="bg-gray-800 p-2">
      <label class="text-sm text-white font-semibold">Vehicle Type:</label>
      <select v-model="selectedType" class="border-none outline-none focus:ring-0 focus:outline-none bg-gray-800 text-white rounded px-2 py-1">
        <option>All</option>
        <option>Truck</option>
        <option>Van</option>
        <option>Electric Car</option>
      </select>
    </div>

    <div>
      <label class="inline-flex items-center">
        <input type="checkbox" v-model="showOnlineOnly" class="form-checkbox text-green-500" />
        <span class="ml-2 text-sm text-white">Online Only</span>
      </label>
    </div>

    <button @click="resetFilters" class="bg-green-500 text-white px-4 py-2 rounded hover:bg-green-600">
      Reset
    </button>

    <button @click="showVehiclesPanel = !showVehiclesPanel; selectedVehicleId = null; markers.value.forEach(m => m.getElement().style.transform = 'scale(1)')" 
            class="ml-auto bg-green-600 text-white px-4 py-2 rounded hover:bg-green-700">
      Vehicles
    </button>
  </div>

  <div ref="mapContainer" class="w-full h-[600px] border shadow-xl"></div>

  <!-- Vehicles Panel -->
  <transition name="slide">
    <div v-if="showVehiclesPanel" class="vehicle-panel">
      <div class="p-4 flex justify-between items-center bg-gray-800 text-green-400 border-b border-gray-700">
        <h2 class="text-lg font-bold">Vehicles</h2>
        <button @click="showVehiclesPanel = false; selectedVehicleId = null; markers.value.forEach(m => m.getElement().style.transform = 'scale(1)')" 
                class="text-gray-400 hover:text-white">✖</button>
      </div>
      <ul class="p-4">
        <li v-for="vehicle in vehicles" :key="vehicle.id" 
    class="vehicle-card cursor-pointer" 
    @click="expandVehicle(vehicle)">
  <!-- Card header -->
  <div class="flex items-center justify-between mb-2">
    <h3 class="text-lg font-bold text-green-400">{{ vehicle.id }}</h3>
    <span :class="{
      'text-green-400': vehicle.status === 'Moving',
      'text-yellow-400': vehicle.status === 'Idle',
      'text-red-400': vehicle.status === 'Stopped'
    }" class="text-xs font-semibold uppercase tracking-wider">{{ vehicle.status }}</span>
  </div>

  <!-- Driver -->
  <div class="text-sm text-gray-400 mb-1">
    Driver: <span class="font-medium text-gray-200">{{ vehicle.driverName }}</span>
  </div>

  <!-- Divider -->
  <div v-if="selectedVehicleId === vehicle.id" class="border-t border-gray-600 mt-3 pt-3">
    <!-- Expanded details -->
    <div class="grid grid-cols-2 gap-x-4 gap-y-2 text-xs text-gray-400">
      <div class="font-semibold">Type</div><div>{{ vehicle.vehicleType }}</div>
      <div class="font-semibold">Online</div><div>{{ vehicle.isOnline ? 'Yes' : 'No' }}</div>
      <div class="font-semibold">Lat</div><div>{{ vehicle.latitude.toFixed(5) }}</div>
      <div class="font-semibold">Lng</div><div>{{ vehicle.longitude.toFixed(5) }}</div>
      <div v-if="vehicle.speed" class="font-semibold">Speed</div>
      <div v-if="vehicle.speed">{{ vehicle.speed }} km/h</div>
    </div>
  </div>
</li>
      </ul>
    </div>
  </transition>
</div>
</template>

<style scoped>
.slide-enter-active, .slide-leave-active {
  transition: all 0.4s cubic-bezier(0.25, 1, 0.5, 1);
}
.slide-enter-from, .slide-leave-to {
  transform: translateX(100%);
  opacity: 0;
}
.vehicle-panel {
  position: absolute;
  top: 64px;
  right: 0;
  width: 300px;
  height: 100%;
  background-color: #1f2937;
  box-shadow: -2px 0 10px rgba(0,0,0,0.2);
  z-index: 50;
  overflow-y: auto;
  border-left: 1px solid #4b5563;
}
.vehicle-card {
  background: #2a2f3a;
  border: 1px solid #3d4451;
  border-radius: 12px;
  padding: 14px 16px;
  margin-bottom: 14px;
  box-shadow: 0 1px 3px rgba(0,0,0,0.2);
  transition: all 0.25s ease-in-out;
}
.vehicle-card:hover {
  transform: translateY(-5px) scale(1.02);
  box-shadow: 0 6px 18px rgba(0,0,0,0.4);
}
</style>
