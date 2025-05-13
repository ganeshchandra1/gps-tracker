<template>
    <div class="h-screen bg-gray-800 text-white flex flex-col transition-all duration-300" :class="{ 'w-16': isCollapsed, 'w-64': !isCollapsed }">
        <button
            class="bg-gray-800 hover:bg-gray-600 text-white p-2 focus:outline-none"
            @click="toggleSidebar"
        >
            <span v-if="isCollapsed">☰</span>
            <span v-else  class="text-white-400 hover:text-white text-2xl font-bold">    <XMarkIcon
                class="h-6 w-6 r-0 text-gray-400 hover:text-white cursor-pointer" /></span>
        </button>
        <ul class="mt-4 space-y-2">
            <li>
                <a
                    href="/map"
                    class="flex items-center space-x-2 px-4 py-2 rounded-md hover:bg-gray-700 transition"
                >
                    <span class="material-icons">map</span>
                    <span v-if="!isCollapsed">Maps</span>
                </a>
            </li>
            <li>
                <a
                    href="#report"
                    class="flex items-center space-x-2 px-4 py-2 rounded-md hover:bg-gray-700 transition"
                >
                    <span class="material-icons">bar_chart</span>
                    <span v-if="!isCollapsed">Report</span>
                </a>
            </li>
            <li>
                <a
                    href="/logout"
                    class="flex items-center space-x-2 px-4 py-2 rounded-md hover:bg-gray-700 transition"
                >
                    <span class="material-icons">person</span>
                    <span v-if="!isCollapsed">Profile</span>
                </a>
            </li>
        </ul>
    </div>
</template>

<script>
import { XMarkIcon } from '@heroicons/vue/24/solid'
export default {
    components: { XMarkIcon },
    data() {
        return {
            isCollapsed: false,
        };
    },
    methods: {
        toggleSidebar() {
            this.isCollapsed = !this.isCollapsed;
        },
    },
};
</script>

<style>
@import url('https://fonts.googleapis.com/icon?family=Material+Icons');

</style>