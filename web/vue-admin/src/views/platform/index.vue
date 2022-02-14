<template>
  <div class="app-container">
    <el-input v-model="search.topic" placeholder="App Id" style="width: 200px;" />
    <el-select v-model="search.status" placeholder="状态" :clearable="clearable" style="margin-left: 10px;">
      <el-option
        v-for="item in optionsStatus"
        :key="item.value"
        :label="item.label"
        :value="item.value"
      />
    </el-select>
    <el-button class="filter-item" style="margin-left: 10px;" type="primary" icon="el-icon-search" @click="handleSearch">
      搜索
    </el-button>
    <el-button class="filter-item" style="margin-left: 10px;" type="primary" icon="el-icon-edit" @click="handleCreate">
      新增
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
      <el-table-column label="名称" width="200px">
        <template slot-scope="scope">
          {{ scope.row.name }}
        </template>
      </el-table-column>
      <el-table-column label="App Id" width="200px">
        <template slot-scope="scope">
          {{ scope.row.app_id }}
        </template>
      </el-table-column>
      <el-table-column label="App Secret" width="auto">
        <template slot-scope="scope">
          {{ scope.row.app_secret }}
        </template>
      </el-table-column>
      <el-table-column label="状态" width="80px">
        <template slot-scope="scope">
          <el-tag v-if="scope.row.status===1" type="success">正常</el-tag>
          <el-tag v-else-if="scope.row.status===0" type="error">禁用</el-tag>
        </template>
      </el-table-column>
      <el-table-column label="备注" width="200px">
        <template slot-scope="scope">
          {{ scope.row.remark }}
        </template>
      </el-table-column>
      <el-table-column label="操作" width="160px">
        <template slot-scope="scope">
          <el-button type="primary" size="mini" @click="reExecute(scope.row.id)">编辑</el-button>
          <el-popconfirm
            confirm-button-text="好的"
            cancel-button-text="不用了"
            icon="el-icon-info"
            icon-color="red"
            title="确定删除吗？"
            @onConfirm="deleteAction(scope.row.id)"
          >
            <el-button slot="reference" type="danger" size="mini" style="margin-left:10px;">
              删除
            </el-button>
          </el-popconfirm>
        </template>
      </el-table-column>
    </el-table>

    <pagination v-show="total>0" :total="total" :page.sync="search.page" @pagination="getList" />

    <el-dialog :title="textMap[dialogStatus]" :visible.sync="dialogFormVisible" width="500px" top="100px">
      <el-form ref="dataForm" :rules="rules" :model="temp" label-position="left" label-width="120px" style="width: 400px; margin-left:50px;">
        <el-form-item label="名称" prop="name">
          <el-input v-model="temp.name" />
        </el-form-item>
        <el-form-item label="备注">
          <el-input v-model="temp.remark" type="textarea" />
        </el-form-item>
        <el-form-item label="是否有效">
          <el-radio-group v-model="temp.status">
            <el-radio :label="1">有效</el-radio>
            <el-radio :label="0">无效</el-radio>
          </el-radio-group>
        </el-form-item>
      </el-form>
      <div slot="footer" class="dialog-footer">
        <el-button @click="dialogFormVisible = false">
          取消
        </el-button>
        <el-button type="primary" @click="dialogStatus==='create'?createData():updateData()">
          确认
        </el-button>
      </div>
    </el-dialog>
  </div>
</template>

<script>
import Pagination from '@/components/Pagination'
import 'vue-json-viewer/style.css'
import { create, getPlatformList, deleteAction } from '@/api/platform'
import { update } from '@/api/consume'
const moment = require('moment')

export default {
  filters: {
    dateFilter(val) {
      return moment(val).format('YYYY-MM-DD hh:mm:ss')
    }
  },
  components: {
    Pagination
  },
  data() {
    return {
      clearable: true,
      optionsStatus: [
        {
          value: 1,
          label: '正常'
        },
        {
          value: 0,
          label: '禁用'
        }
      ],
      search: {
        AppId: null,
        status: null,
        page: 1
      },
      rules: {
        topic: [{ required: true, message: '必要选项', trigger: 'blur' }],
        channel: [{ required: true, message: '必要选项', trigger: 'blur' }]
      },
      temp: {
        name: '',
        remark: '',
        status: 1 // 1有效 0无效
      },
      textMap: {
        update: '修改',
        create: '新增'
      },
      dialogStatus: '',
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
    getList() {
      this.listLoading = true
      getPlatformList({ ...this.search }).then(response => {
        this.listLoading = false
        this.list = response.result.result
        this.search.page = response.result.page
        this.total = response.result.total
      })
    },
    handleCreate() {
      this.resetTemp()
      this.dialogStatus = 'create'
      this.dialogFormVisible = true
      this.$nextTick(() => {
        this.$refs['dataForm'].clearValidate()
      })
    },
    resetTemp() {
      this.temp = {
        name: '',
        remark: '',
        status: 1 // 0有效 1无效
      }
    },
    updateData() {
      this.$refs['dataForm'].validate((valid) => {
        if (valid) {
          const tempData = Object.assign({}, this.temp)
          update(tempData).then(() => {
            this.dialogFormVisible = false
            this.$notify({
              title: 'Success',
              message: '更新成功',
              type: 'success',
              duration: 2000
            })
            this.getList()
          })
        }
      })
    },
    createData() {
      this.$refs['dataForm'].validate((valid) => {
        if (valid) {
          console.log(JSON.stringify(this.temp))
          create(this.temp).then(() => {
            this.dialogFormVisible = false
            this.$notify({
              title: '新增成功',
              message: '新增成功',
              type: 'success',
              duration: 2000
            })
            this.getList()
          })
        }
      })
    },
    deleteAction(id) {
      deleteAction({ id }).then(() => {
        this.dialogFormVisible = false
        this.$notify({
          title: '删除成功',
          message: '删除成功',
          type: 'success',
          duration: 2000
        })
        this.getList()
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
