<script setup>
import { ref, onMounted } from 'vue'
const API_BASE_URL = import.meta.env.VITE_BACKEND_URL
const vehicles = ref([])
const alerts = ref([])
const loading = ref(true)
const lastUpdated = ref(null)

const totalVehicles = ref(0)
const onlineVehicles = ref(0)
const offlineVehicles = ref(0)
const movingVehicles = ref(0)
const idleVehicles = ref(0)
const stoppedVehicles = ref(0)
const maintenanceDueCount = ref(0)
const averageEngineHours = ref(0)
const vehicleTypeCounts = ref({})
const recentActivity = ref([])

const fetchVehicles = async () => {
  try {
    const response = await fetch(`${API_BASE_URL}vehicles`)
    const data = await response.json()
    vehicles.value = data
    computeStats()
    lastUpdated.value = new Date().toLocaleString()
  } catch (error) {
    console.error('Error fetching vehicles:', error)
  } finally {
    loading.value = false
  }
}

const fetchAlerts = async () => {
  try {
    const response = await fetch(`${API_BASE_URL}alerts`)
    const data = await response.json()
    alerts.value = data
  } catch (error) {
    console.error('Error fetching alerts:', error)
  }
}

const computeStats = () => {
  totalVehicles.value = vehicles.value.length
  onlineVehicles.value = vehicles.value.filter(v => v.isOnline).length
  offlineVehicles.value = totalVehicles.value - onlineVehicles.value
  movingVehicles.value = vehicles.value.filter(v => v.status === 'Moving').length
  idleVehicles.value = vehicles.value.filter(v => v.status === 'Idle').length
  stoppedVehicles.value = vehicles.value.filter(v => v.status === 'Stopped').length
  maintenanceDueCount.value = vehicles.value.filter(v => v.maintenanceDue).length

  const totalEngineHours = vehicles.value.reduce((sum, v) => sum + v.engineHours, 0)
  averageEngineHours.value = totalVehicles.value ? (totalEngineHours / totalVehicles.value).toFixed(1) : 0

  const typeCounts = {}
  vehicles.value.forEach(v => {
    typeCounts[v.vehicleType] = (typeCounts[v.vehicleType] || 0) + 1
  })
  vehicleTypeCounts.value = typeCounts

  recentActivity.value = [...vehicles.value]
    .sort((a, b) => new Date(b.lastUpdate) - new Date(a.lastUpdate))
    .slice(0, 5)
}

const refreshData = () => {
  loading.value = true
  fetchVehicles()
  fetchAlerts()
}

onMounted(() => {
  fetchVehicles()
  fetchAlerts()
})
</script>

<template>
  <div class="min-h-screen bg-gray-900 text-white">
    <div class="max-w-7xl mx-auto px-4 sm:px-6 lg:px-8 py-6 space-y-8">
      <!-- Header -->
      <div class="flex justify-between items-center mb-6">
        <h1 class="text-3xl font-bold text-green-400">Fleet Dashboard</h1>
        <div class="flex items-center space-x-4">
          <span class="text-sm text-gray-400">Last Updated: {{ lastUpdated }}</span>
          <button @click="refreshData" class="bg-green-500 hover:bg-green-600 text-white px-3 py-1 rounded">
            Refresh
          </button>
        </div>
      </div>

      <!-- Loading State -->
      <div v-if="loading" class="flex justify-center items-center h-64">
        <svg class="animate-spin h-10 w-10 text-green-400" xmlns="http://www.w3.org/2000/svg" fill="none" viewBox="0 0 24 24">
          <circle class="opacity-25" cx="12" cy="12" r="10" stroke="currentColor" stroke-width="4"></circle>
          <path class="opacity-75" fill="currentColor" d="M4 12a8 8 0 018-8v8H4z"></path>
        </svg>
      </div>

      <!-- Dashboard Content -->
      <div v-else class="space-y-8">
        <!-- Analytics Cards -->
        <div class="grid grid-cols-1 sm:grid-cols-2 lg:grid-cols-3 gap-6">
          <div class="bg-gray-800 p-5 rounded-xl shadow hover:shadow-lg transition transform hover:scale-105">
            <h2 class="text-sm text-gray-400 uppercase mb-2">Total Vehicles</h2>
            <p class="text-3xl font-bold text-green-400">{{ totalVehicles }}</p>
          </div>
          <div class="bg-gray-800 p-5 rounded-xl shadow hover:shadow-lg transition transform hover:scale-105">
            <h2 class="text-sm text-gray-400 uppercase mb-2">Online</h2>
            <p class="text-3xl font-bold text-blue-400">{{ onlineVehicles }}</p>
          </div>
          <div class="bg-gray-800 p-5 rounded-xl shadow hover:shadow-lg transition transform hover:scale-105">
            <h2 class="text-sm text-gray-400 uppercase mb-2">Offline</h2>
            <p class="text-3xl font-bold text-red-400">{{ offlineVehicles }}</p>
          </div>
          <div class="bg-gray-800 p-5 rounded-xl shadow hover:shadow-lg transition transform hover:scale-105">
            <h2 class="text-sm text-gray-400 uppercase mb-2">Moving</h2>
            <p class="text-3xl font-bold text-yellow-300">{{ movingVehicles }}</p>
          </div>
          <div class="bg-gray-800 p-5 rounded-xl shadow hover:shadow-lg transition transform hover:scale-105">
            <h2 class="text-sm text-gray-400 uppercase mb-2">Idle</h2>
            <p class="text-3xl font-bold text-purple-300">{{ idleVehicles }}</p>
          </div>
          <div class="bg-gray-800 p-5 rounded-xl shadow hover:shadow-lg transition transform hover:scale-105">
            <h2 class="text-sm text-gray-400 uppercase mb-2">Stopped</h2>
            <p class="text-3xl font-bold text-orange-300">{{ stoppedVehicles }}</p>
          </div>
        </div>

        <!-- Fleet Insights -->
        <div class="grid grid-cols-1 md:grid-cols-3 gap-6">
          <div class="bg-gray-800 p-5 rounded-xl shadow hover:shadow-lg transition transform hover:scale-105">
            <h2 class="text-sm text-gray-400 uppercase mb-2">Maintenance Due</h2>
            <p class="text-3xl font-bold text-red-400">{{ maintenanceDueCount }}</p>
          </div>
          <div class="bg-gray-800 p-5 rounded-xl shadow hover:shadow-lg transition transform hover:scale-105">
            <h2 class="text-sm text-gray-400 uppercase mb-2">Avg. Engine Hours</h2>
            <p class="text-3xl font-bold text-cyan-300">{{ averageEngineHours }}</p>
          </div>
          <div class="bg-gray-800 p-5 rounded-xl shadow hover:shadow-lg transition transform hover:scale-105">
            <h2 class="text-sm text-gray-400 uppercase mb-2">Vehicle Types</h2>
            <ul class="mt-2 space-y-1">
              <li v-for="(count, type) in vehicleTypeCounts" :key="type" class="flex justify-between text-sm text-white">
                <span>{{ type }}</span>
                <span class="font-semibold text-green-300">{{ count }}</span>
              </li>
            </ul>
          </div>
        </div>

        <!-- Recent Alerts -->
        <div class="bg-gray-800 p-6 rounded-xl shadow hover:shadow-lg transition transform hover:scale-105">
          <h2 class="text-lg font-semibold text-white mb-4">Recent Alerts</h2>
          <ul class="space-y-3">
            <li v-for="alert in alerts" :key="alert.timestamp" class="bg-gray-700 p-4 rounded-lg">
              <div class="flex justify-between items-center">
                <div>
                  <div class="text-sm text-gray-400">Vehicle ID: {{ alert.vehicleId }}</div>
                  <div class="text-white font-semibold">{{ alert.message }}</div>
                </div>
                <div class="text-sm text-gray-400">{{ new Date(alert.timestamp).toLocaleTimeString() }}</div>
              </div>
            </li>
          </ul>
        </div>

        <!-- Recent Vehicle Activity -->
        <div class="bg-gray-800 p-6 rounded-xl shadow hover:shadow-lg transition transform hover:scale-105">
          <h2 class="text-lg font-semibold text-white mb-4">Recent Vehicle Activity</h2>
          <ul class="space-y-3">
            <li v-for="vehicle in recentActivity" :key="vehicle.id" class="bg-gray-700 p-4 rounded-lg">
              <div class="flex justify-between items-center">
                <div>
                  <div class="text-sm text-gray-400">Driver: {{ vehicle.driverName }}</div>
                  <div class="text-white font-semibold">Status: {{ vehicle.status }}</div>
                </div>
                <div class="text-sm text-gray-400">{{ new Date(vehicle.lastUpdate).toLocaleTimeString() }}</div>
              </div>
            </li>
          </ul>
        </div>
      </div>
    </div>
  </div>
</template>

<style scoped>
/* Add any additional custom styles here if needed */
</style>
