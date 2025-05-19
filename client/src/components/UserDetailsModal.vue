<template>
  <a-modal :open="visible" title="User Details" @ok="handleSubmit" @cancel="handleCancel">
    <a-descriptions bordered :column="3">
      <a-descriptions-item label="Id" :span="3">
        {{ user?.id }}
      </a-descriptions-item>
      <a-descriptions-item label="Name" :span="3">
        <input label="Name" v-model="updateName"
      />
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
import type { User } from '@/types/user'
import { ref, watch } from 'vue';

const { visible, user } = defineProps<{
  visible: boolean
  user: User | null
}>()

const updateName = ref<string>(user?.name ?? '')

const emit = defineEmits<{
  (e: 'update:visible', value: boolean): void
  (e: 'update', value: User): void
}>()

const handleCancel = () => {
  emit('update:visible', false)
}

const handleSubmit = (): void => {
  const updateUser: User = {
    id: user?.id ?? '',
    name: updateName.value ?? '',
    email: user?.email ?? '',
    password: user?.password ?? '',
    updated_at: user?.updated_at ?? '',
    created_at: user?.created_at ?? '',
  }

  emit('update', updateUser)
}

watch(
  () => user,
  (newUser: User | null) => {
    updateName.value = newUser?.name ?? ''
  },
)

</script>
