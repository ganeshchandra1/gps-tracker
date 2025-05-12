<script setup>
import { RouterLink, RouterView } from 'vue-router'
import Navbar from './components/Navbar.vue'
import SideBar from './components/SideBar.vue'
</script>

<template>
<div class="flex flex-col w-full min-h-screen">
  <Navbar />
  <div class="flex flex-row w-full">
    <SideBar />
    <RouterView />
</div>
</div>
</template>

<style scoped>

</style>
