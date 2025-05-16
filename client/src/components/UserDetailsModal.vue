<template>
  <a-modal :visible="visible" title="User Details" @cancel="handleCancel" :footer="null">
    <a-descriptions bordered>
      <a-descriptions-item label="Name" :span="3">
        {{ user?.name }}
      </a-descriptions-item>
      <a-descriptions-item label="Email" :span="3">
        {{ user?.email }}
      </a-descriptions-item>
      <a-descriptions-item label="Password" :span="3">
        {{ user?.password }}
      </a-descriptions-item>
      <a-descriptions-item label="Created At" :span="3">
        {{ user?.created_at }}
      </a-descriptions-item>
      <a-descriptions-item label="Updated At" :span="3">
        {{ user?.updated_at }}
      </a-descriptions-item>
    </a-descriptions>
  </a-modal>
</template>

<script lang="ts" setup>
import type { User } from '@/stores/user'

const props = defineProps<{
  visible: boolean
  user: User | null
}>()

const emit = defineEmits<{
  (e: 'update:visible', value: boolean): void
}>()

const handleCancel = () => {
  emit('update:visible', false)
}
</script>
