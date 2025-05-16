<template>
  <a-modal :open="visible" title="User Details" @ok="handleOk" @cancel="handleCancel">
    <template v-if="user">
      <p><strong>Id:</strong> {{ user.id }}</p>
      <p><strong>Name:</strong> {{ user.name }}</p>
      <p><strong>Email:</strong> {{ user.email }}</p>
      <p><strong>Password:</strong> {{ user.password }}</p>
      <p><strong>Created At:</strong> {{ user.created_at }}</p>
      <p><strong>Updated At:</strong> {{ user.updated_at }}</p>
    </template>
  </a-modal>
</template>

<script lang="ts" setup>
import type { User } from '@/stores/user'

interface Props {
  visible: boolean
  user: User | null
}

interface Emits {
  (e: 'update:visible', value: boolean): void
}

const props = defineProps<Props>()
const emit = defineEmits<Emits>()

const handleOk = () => {
  emit('update:visible', false)
}

const handleCancel = () => {
  emit('update:visible', false)
}
</script>
