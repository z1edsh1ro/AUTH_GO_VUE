<template>
  <a-layout>
    <Navbar />
    <a-layout-content>
      <HeroSectionComponent />
    </a-layout-content>
  </a-layout>
</template>

<script lang="ts" setup>
import HeroSectionComponent from '@/components/HeroSection.vue'
import Navbar from '@/components/Navbar.vue'
</script>
