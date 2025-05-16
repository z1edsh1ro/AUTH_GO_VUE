<template>
  <div class="page-container">
    <Navbar />
    <div class="content">
      <a-spin :spinning="userStore.loading">
        <a-table :columns="columns" :data-source="userStore.users">
          <template #bodyCell="{ column, record }">
            <template v-if="column.key === 'operation'">
              <a @click="showModal(record)">View Details</a>
            </template>
          </template>
        </a-table>
      </a-spin>

      <UserDetailsModal
        :visible="isModalVisible"
        @update:visible="isModalVisible = $event"
        :user="selectedUser"
      />

      <a-alert
        v-if="userStore.error"
        type="error"
        :message="userStore.error"
        show-icon
      />
    </div>
  </div>
</template>

<script lang="ts" setup>
import { ref, onMounted } from 'vue'
import { useRouter } from 'vue-router'
import type { TableColumnsType } from 'ant-design-vue'
import { useUserStore } from '@/stores/user'
import { useAuthStore } from '@/stores/auth'
import type { User } from '@/stores/user'
import UserDetailsModal from '@/components/UserDetailsModal.vue'
import Navbar from '@/components/Navbar.vue'

const router = useRouter()
const userStore = useUserStore()
const authStore = useAuthStore()
const isModalVisible = ref(false)
const selectedUser = ref<User | null>(null)

// Column definitions
const columns = ref<TableColumnsType<User>>([
  {
    title: 'Name',
    dataIndex: 'name',
    key: 'name',
  },
  {
    title: 'Email',
    dataIndex: 'email',
    key: 'email',
  },
  {
    title: 'Password',
    dataIndex: 'password',
    key: 'password',
  },
  {
    title: 'createdAt',
    dataIndex: 'created_at',
    key: 'created_at',
  },
  {
    title: 'updatedAt',
    dataIndex: 'updated_at',
    key: 'updated_at',
  },
  {
    title: 'Action',
    key: 'operation',
  },
])

const showModal = (record: User) => {
  selectedUser.value = record
  isModalVisible.value = true
}

onMounted(() => {
  if (!authStore.isAuthenticated) {
    router.push('/login')
    return
  }
  userStore.fetchUsers()
})
</script>

<style scoped>
.page-container {
  min-height: 100vh;
  background: #f0f2f5;
}

.content {
  padding: 24px;
  margin-top: 64px;
  min-height: calc(100vh - 64px);
  background: #fff;
}
</style>
