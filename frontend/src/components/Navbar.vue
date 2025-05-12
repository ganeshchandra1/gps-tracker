<template>
    <nav class="bg-gray-800 text-white px-6 py-4 shadow-md">
      <div class="container mx-auto flex items-center justify-between">
        <!-- Brand -->
        <router-link to="/" class="text-2xl font-bold text-green-500 hover:text-gray-300">
          Fleet Tracker
        </router-link>
        <div class="flex items-center space-x-4">
          <router-link 
            to="/logout" 
            class="hover:text-gray-300"
            :class="{ 'font-bold text-green-500 border-white': $route.path.startsWith('/logout') }"
          >
            Log out
          </router-link>
  
          <router-link 
            to="/profile" 
            class="hover:text-gray-300 flex items-center"
            :class="{ 'font-bold  border-white': $route.path.startsWith('/profile') }"
          >
            <UserCircleIcon class="h-5 w-5" />
          </router-link>
        </div>
        </div>

    </nav>
  </template>
  
  <script setup>
  import { UserCircleIcon } from '@heroicons/vue/24/solid'
  </script>
  