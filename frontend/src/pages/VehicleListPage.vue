<script setup>
import { ref, onMounted } from 'vue'
import { useRouter } from 'vue-router'

const vehicles = ref([])
const router = useRouter()

onMounted(async () => {
  const response = await fetch('http://localhost:8080/api/vehicles')
  vehicles.value = await response.json()
})

const goToDetails = (id) => {
  router.push(`/vehicles/${id}`)
}
</script>

<template>
<div class="min-h-screen bg-gray-100 dark:bg-gray-900 p-6">
  <h1 class="text-3xl font-extrabold text-center text-green-600 mb-8 animate-fade-in">Fleet Vehicles</h1>

  <div class="grid grid-cols-1 sm:grid-cols-2 md:grid-cols-3 lg:grid-cols-4 gap-6 animate-fade-in">
    <div 
      v-for="vehicle in vehicles" 
      :key="vehicle.id"
      @click="goToDetails(vehicle.id)"
      class="bg-white dark:bg-gray-800 rounded-xl shadow-md hover:shadow-xl transform hover:scale-105 transition duration-300 cursor-pointer overflow-hidden"
    >
      <div class="p-5">
        <h2 class="text-xl font-bold text-gray-800 dark:text-gray-100">{{ vehicle.id }}</h2>
        <p class="text-sm text-gray-500 dark:text-gray-400 mb-2">{{ vehicle.vehicleType }}</p>

        <div class="mb-4 grid grid-cols-2 gap-1 text-xs text-gray-600 dark:text-gray-300">
          <div>Driver:</div><div>{{ vehicle.driverName }}</div>
          <div>Odometer:</div><div>{{ vehicle.odometer }} km</div>
          <div>Last Service:</div><div>{{ new Date(vehicle.lastServiceDate).toLocaleDateString() }}</div>
          <div>Engine Hours:</div><div>{{ vehicle.engineHours }}</div>
          <div>{{ vehicle.vehicleType === 'Electric Car' ? 'Battery' : 'Fuel' }}:</div>
          <div>{{ vehicle.vehicleType === 'Electric Car' ? vehicle.batteryStatus + '%' : vehicle.fuelLevel + '%' }}</div>
        </div>

        <div class="mt-3 flex items-center justify-between">
          <span 
            :class="[
              'px-2 py-1 text-xs font-semibold rounded',
              vehicle.status === 'Moving' ? 'bg-green-200 text-green-800 dark:bg-green-700 dark:text-green-200' :
              vehicle.status === 'Idle' ? 'bg-yellow-200 text-yellow-800 dark:bg-yellow-700 dark:text-yellow-200' :
              'bg-red-200 text-red-800 dark:bg-red-700 dark:text-red-200'
            ]"
          >{{ vehicle.status }}</span>

          <span>
            <svg v-if="vehicle.isOnline" class="h-5 w-5 text-green-500" fill="currentColor" viewBox="0 0 20 20"><path fill-rule="evenodd" d="M10 18a8 8 0 100-16 8 8 0 000 16zm3.707-9.707a1 1 0 00-1.414-1.414L9 10.586 7.707 9.293a1 1 0 00-1.414 1.414L9 13.414l4.707-4.707z" clip-rule="evenodd"/></svg>
            <svg v-else class="h-5 w-5 text-red-500" fill="currentColor" viewBox="0 0 20 20"><path fill-rule="evenodd" d="M10 18a8 8 0 100-16 8 8 0 000 16zm-2-9a1 1 0 112 0v2a1 1 0 11-2 0v-2zm1-4a1 1 0 100 2 1 1 0 000-2z" clip-rule="evenodd"/></svg>
          </span>
        </div>
      </div>
    </div>
  </div>
</div>
</template>

<style>
@keyframes fade-in {
  from { opacity: 0; transform: translateY(10px); }
  to { opacity: 1; transform: translateY(0); }
}

.animate-fade-in {
  animation: fade-in 0.8s ease-out;
}
</style>
