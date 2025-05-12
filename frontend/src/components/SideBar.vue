<template>
    <div class="h-screen bg-gray-800 text-white flex flex-col transition-all duration-300" :class="{ 'w-16': isCollapsed, 'w-64': !isCollapsed }">
        <button
            class="bg-gray-700 hover:bg-gray-600 text-white p-2 focus:outline-none"
            @click="toggleSidebar"
        >
            <span v-if="isCollapsed">☰</span>
            <span v-else>><</span>
        </button>
        <ul class="mt-4 space-y-2">
            <li>
                <a
                    href="/"
                    class="flex items-center space-x-2 px-4 py-2 rounded-md hover:bg-gray-700 transition"
                >
                    <span class="material-icons">dashboard</span>
                    <span v-if="!isCollapsed">Dashboard</span>
                </a>
            </li>
            <li>
                <a
                    href="/vehicles"
                    class="flex items-center space-x-2 px-4 py-2 rounded-md hover:bg-gray-700 transition"
                >
                    <span class="material-icons">directions_car</span>
                    <span v-if="!isCollapsed">Vehicles</span>
                </a>
            </li>
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
                    href="#profile"
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
export default {
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