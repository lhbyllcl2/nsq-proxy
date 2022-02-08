<template>
  <div class="app-container">
    <el-input v-model="search.topic" placeholder="Topic" style="width: 200px;" />
    <el-select v-model="search.status" placeholder="执行状态" :clearable="clearable" style="margin-left: 10px;">
      <el-option
        v-for="item in optionsStatus"
        :key="item.value"
        :label="item.label"
        :value="item.value"
      />
    </el-select>
    <el-select v-model="search.msgType" placeholder="消息类型" :clearable="clearable" style="margin-left: 10px;">
      <el-option
        v-for="item in typeOption"
        :key="item.value"
        :label="item.label"
        :value="item.value"
      />
    </el-select>
    <el-button class="filter-item" style="margin-left: 10px;" type="primary" icon="el-icon-search" @click="handleSearch">
      搜索
    </el-button>
    <el-table
      v-loading="listLoading"
      :data="list"
      element-loading-text="Loading"
      border
      fit
      highlight-current-row
      class="common-table"
    >
      <el-table-column label="ID" width="60px">
        <template slot-scope="scope">
          {{ scope.row.id }}
        </template>
      </el-table-column>
      <el-table-column label="消息Id" width="140px">
        <template slot-scope="scope">
          {{ scope.row.message_id }}
        </template>
      </el-table-column>
      <el-table-column label="Topic" width="100px">
        <template slot-scope="scope">
          {{ scope.row.topic }}
        </template>
      </el-table-column>
      <el-table-column label="请求地址" width="280px">
        <template slot-scope="scope">
          {{ scope.row.url }}
        </template>
      </el-table-column>
      <el-table-column label="请求方法" width="100px">
        <template slot-scope="scope">
          {{ scope.row.method.toUpperCase() }}
        </template>
      </el-table-column>
      <el-table-column label="请求参数">
        <template slot-scope="scope">
          <a v-if="scope.row.argument!==''" @click="showParams(scope.row.argument)">点击查看</a>
          <span v-else>--</span>
        </template>
      </el-table-column>
      <el-table-column label="延时(秒)" width="100px">
        <template slot-scope="scope">
          <span>{{ scope.row.delay }}</span>
        </template>
      </el-table-column>
      <el-table-column label="添加时间" width="160px">
        <template slot-scope="scope">{{ scope.row.create_at | dateFilter }}</template>
      </el-table-column>
      <el-table-column label="执行时间" width="160px">
        <template slot-scope="scope">{{ scope.row.update_at | dateFilter }}</template>
      </el-table-column>
      <el-table-column label="状态" width="80px">
        <template slot-scope="scope">
          <el-tag v-if="scope.row.status===1" type="success">成功</el-tag>
          <el-tag v-else-if="scope.row.status===0">待执行</el-tag>
          <el-tag v-else type="danger">失败</el-tag>
        </template>
      </el-table-column>
      <el-table-column label="操作" width="120px">
        <template slot-scope="scope">
          <el-button v-if="scope.row.status===2" type="primary" size="mini" @click="reExecute(scope.row.id)">重新执行</el-button>
          <span v-else>--</span>
        </template>
      </el-table-column>
    </el-table>

    <pagination v-show="total>0" :total="total" :page.sync="search.page" @pagination="getList" />

    <el-dialog :visible.sync="dialogFormVisible" width="500px" top="100px">
      <json-viewer
        :value="jsonData"
        :expand-depth="5"
        copyable
        boxed
        sort
      />
    </el-dialog>
  </div>
</template>

<script>
import { getList } from '@/api/message'
import Pagination from '@/components/Pagination'
import JsonViewer from 'vue-json-viewer/ssr'
import 'vue-json-viewer/style.css'
const moment = require('moment')

export default {
  filters: {
    dateFilter(val) {
      return moment(val).format('YYYY-MM-DD hh:mm:ss')
    }
  },
  components: {
    Pagination,
    JsonViewer
  },
  data() {
    return {
      clearable: true,
      jsonData: null,
      typeOption: [
        {
          value: 1,
          label: '实时消息'
        },
        {
          value: 2,
          label: '延迟消息'
        }
      ],
      optionsStatus: [
        {
          value: 0,
          label: '待执行'
        },
        {
          value: 1,
          label: '成功'
        },
        {
          value: 2,
          label: '失败'
        }
      ],
      search: {
        topic: null,
        status: null,
        msgType: null,
        page: 1
      },
      total: 0,
      list: [],
      listLoading: false,
      dialogFormVisible: false
    }
  },
  created() {
    this.getList()
  },
  methods: {
    handleSearch() {
      this.getList()
    },
    reExecute() {
      return ''
    },
    showParams(jsonData) {
      this.jsonData = JSON.parse(jsonData)
      this.dialogFormVisible = true
    },
    getList() {
      this.listLoading = true
      getList({ ...this.search }).then(response => {
        this.listLoading = false
        this.list = response.result.result
        this.search.page = response.result.page
        this.total = response.result.total
      })
    }
  }
}
</script>
<style lang="scss" scoped>
.common-table{
  margin-top: 10px;
}
</style>
