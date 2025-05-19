<template>
  <a-layout>
    <Navbar />
    <a-layout-content>
      <a-row>
        <a-col :span="24">
          <a-card>
            <template>
              <a-space>
                <span>User Management</span>
              </a-space>
            </template>

            <a-table :columns="columns" :data-source="userStore.users">
              <template #bodyCell="{ column, record }">
                <template v-if="column.key === 'operation'">
                  <a @click="showModal(record)"><EyeOutlined /></a>
                </template>
              </template>
            </a-table>
          </a-card>
        </a-col>
      </a-row>

      <UserDetailsModal v-model:visible="isModalVisible" :user="selectedUser" @update="onUpdate" />
    </a-layout-content>
  </a-layout>
</template>

<script lang="ts" setup>
import { ref, onMounted } from 'vue'
import type { TableColumnsType } from 'ant-design-vue'
import type { User } from '@/types/user'
import UserDetailsModal from '@/components/UserDetailsModal.vue'
import Navbar from '@/components/Navbar.vue'
import EyeOutlined from '@ant-design/icons-vue/EyeOutlined'
import { useUserStore } from '@/stores/user'

const userStore = useUserStore()
const isModalVisible = ref<boolean>(false)
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

const showModal = (record: User):void => {
  selectedUser.value = record
  isModalVisible.value = true
}

const onUpdate = (user: User):void => {
  userStore.updateUser(user)
  isModalVisible.value = false
}

onMounted(() => {
  userStore.getUsers()
})
</script>
