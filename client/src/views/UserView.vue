<template>
  <a-layout>
    <Navbar />
    <a-layout-content>
      <a-row>
        <a-col :span="24">
          <a-card>
            <template #title>
              <a-space>
                <span>User Management</span>
              </a-space>
            </template>

            <a-table :columns="columns" :data-source="userStore.users" :loading="userStore.loading">
              <template #bodyCell="{ column, record }">
                <template v-if="column.key === 'operation'">
                  <a @click="showModal(record)"><EyeOutlined /></a>
                </template>
              </template>
            </a-table>

            <a-alert
              v-if="userStore.error"
              type="error"
              :message="userStore.error"
              show-icon
              banner
            />
          </a-card>
        </a-col>
      </a-row>

      <UserDetailsModal
        :visible="isModalVisible"
        @update:visible="isModalVisible = $event"
        :user="selectedUser"
      />
    </a-layout-content>
  </a-layout>
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
import EyeOutlined from '@ant-design/icons-vue/EyeOutlined'

const router = useRouter()
const userStore = useUserStore()
const authStore = useAuthStore()
const isModalVisible = ref(false)
const selectedUser = ref<User | null>(null)

const columns = ref<TableColumnsType<User>>([
  {
    title: 'Id',
    dataIndex: 'id',
    key: 'id',
  },
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
.ant-layout-content {
  padding: 24px;
  min-height: calc(100vh - 64px);
}
</style>
