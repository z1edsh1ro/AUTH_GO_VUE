<template>
  <a-layout>
    <Navbar />
    <a-layout-content class="container">
      <a-col :span="24">
        <a-typography-title :level="1" class="hero-title"> Welcome App </a-typography-title>
      </a-col>
    </a-layout-content>
  </a-layout>
</template>

<script lang="ts" setup>
import Navbar from '@/components/Navbar.vue'
</script>

<style scoped>
.container {
  text-align: center;
  padding: 80px 24px;
  background: linear-gradient(135deg, #1890ff 0%, #722ed1 100%);
}

.hero-title {
  color: white;
  margin-bottom: 24px;
}
</style>
