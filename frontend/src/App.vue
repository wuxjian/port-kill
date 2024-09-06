<script setup>
import { onMounted, ref, reactive } from 'vue'
import { GetProcessInfoByPort, Kill } from '../wailsjs/go/main/App'


onMounted(() => {
  handleQuery()
})

async function handleQuery() {
  if (!port.value) {
    return
  }
  try {
    const result = await GetProcessInfoByPort(port.value + '')
    result.forEach(item => item.key = item.local_address)
    data.value = result
  } catch (e) {
    data.value = []
    console.log('查询端口信息失败', e)
  }
}

async function handleKill() {
  if (!selectedKeys.value.length) {
    return
  }
  visible.value = true
}

const port = ref(null)

const columns = [
  {
    title: '协议',
    dataIndex: 'protocol',
  },
  {
    title: 'IP地址及端口',
    dataIndex: 'local_address',
  },
  {
    title: '状态',
    dataIndex: 'state',
  },
  {
    title: '进程ID',
    dataIndex: 'pid',
  },
];

const rowSelection = reactive({
  type: 'checkbox',
  showCheckedAll: false,
  onlyCurrent: true,
});
const selectedKeys = ref([]);

const data = ref([]);

const visible = ref(false);
const handleOk = async () => {
  // 循环selectedKeys
  if (!selectedKeys.value.length) {
    return
  }
  const pidArray = data.value.filter(item => {
    return selectedKeys.value.includes(item.local_address)
  }).map(item => item.pid)
  try {
    await Kill([...new Set(pidArray)].join(','))
    handleQuery()
  } catch (e) {
    console.log('error', e)
  }
  visible.value = false;
};
const handleCancel = () => {
  visible.value = false;
}
</script>

<template>
  <div>
    <a-space direction="vertical" fill>
      <div>
        <a-space direction="horizontal" fill>
          <a-input-number v-model="port" :style="{ width: '320px' }" placeholder="端口号" :min="10" :max="65535" />
          <a-button type="primary" status="success" @click="handleQuery">查询</a-button>
          <a-button type="primary" status="warning" @click="handleKill">Kill</a-button>
        </a-space>
      </div>

      <a-table :columns="columns" :data="data" :pagination="false" :row-selection="rowSelection"
        v-model:selectedKeys="selectedKeys" />
    </a-space>
  </div>
  <a-modal v-model:visible="visible" @ok="handleOk" @cancel="handleCancel">
    <template #title>
      确认
    </template>
    <div>确认要删除进程?</div>
  </a-modal>

</template>

<style></style>
