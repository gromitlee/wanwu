<template>
  <div class="depart-all">
    <el-select
      ref="selectTree"
      v-model="valueName"
      multiple
      clearable
      @remove-tag="changeValue"
      @clear="clearHandle"
      :disabled="disabled"
      :placeholder="$t('common.select.placeholder')"
      style="width: 100%"
    >
      <div class="dropdown-tree-wrap" style="padding: 8px 14px">
        <div style="padding-left: 22px; padding-bottom: 10px" v-if="data.length">
          <el-checkbox @change="handleAllChange" v-model="allChecked">{{$t('role.tree.all')}}</el-checkbox>
        </div>
        <el-option :value="valueName" style="height: auto; background: #fff">
          <el-tree
            :props="{...defaultProps, ...treeKeyMap}"
            :data="data"
            show-checkbox
            ref="tree"
            :node-key="treeKeyMap.value || defaultProps.value"
            id="tree-option"
            default-expand-all
            @check-change="handleNodeClick"
            style="margin: -1px -40px -1px -20px;"
          />
        </el-option>
      </div>
    </el-select>
  </div>
</template>

<script>
export default {
  name: 'selectTree',
  props: {
    dataArray: [],
    defaultValue: [],
    dataList: [],
    treeKeyMap: {},
    disabled: false
  },
  data () {
    return {
      allChecked: false,
      valueName: [],
      choosedValue: [],
      data: [],
      defaultProps: {
        value: "id",
        children: "children",
        label: "name",
      },
    }
  },
  watch: {
    defaultValue(newVal) {
      const {value, label} = this.defaultProps
      this.choosedValue = newVal.map(item => item[value])
      this.valueName = newVal.map(item => item[label])
      this.$refs.tree.setCheckedKeys(this.choosedValue)
    }
  },
  created() {},
  mounted() {
    this.defaultProps = {...this.defaultProps, ...this.treeKeyMap}
    this.data = this.dataList || []

    const {value, label} = this.defaultProps
    this.choosedValue = this.defaultValue.map(item => item[value])
    this.valueName = this.defaultValue.map(item => item[label])

    this.$refs.tree.setCheckedKeys(this.choosedValue)
  },
  methods: {
    // 全选
    handleAllChange(val){
      if (val) {
        const findChild = (data) => {
          const {label, value, children} = this.defaultProps
          data.forEach(item => {
            if (!(item[children] && item[children].length)) {
              if (!this.choosedValue.includes(item[value]))
                this.choosedValue.push(item[value])
              if (!this.valueName.includes(item[label]))
                this.valueName.push(item[label])
            } else {
              findChild(item[children])
            }
          })
        }
        findChild(this.data)
        this.choosedValue = [...new Set(this.choosedValue)]
        this.valueName = [...new Set(this.valueName)]
        this.$refs.tree.setCheckedKeys(this.choosedValue)
        this.allChecked = true
      } else {
        this.allChecked = false
        this.choosedValue = []
        this.$refs.tree.setCheckedKeys([])
      }
      // this.$emit('handleChange', this.choosedValue)
    },
    // 点击树形选择节点
    handleNodeClick(node, checked) {
      const {value, label, children} = this.defaultProps

      if (checked) {
        if (!node[children]) {
          const hasName = this.valueName.includes(node[label])
          if (!hasName) {
            this.valueName.push(node[label])
            this.choosedValue.push(node[value])
            this.choosedValue = [...new Set(this.choosedValue)]
            this.valueName = [...new Set(this.valueName)]
          }
        }
      } else {
        this.allChecked = false
        if (!node[children]) {
          const nameIndex = this.valueName.findIndex(item => item === node[label])
          if (nameIndex >= 0) this.valueName.splice(nameIndex, 1)

          const valueIndex = this.choosedValue.findIndex(item => item === node[value])
          if (valueIndex >= 0) this.choosedValue.splice(valueIndex, 1)
        }
      }
      this.$emit('handleChange', this.choosedValue)
    },
    // 删除单个标签
    changeValue(val) {
      this.allChecked = false
      const curObj = this.findItemByName(this.data, val)
      const index = this.choosedValue.findIndex(item => item === curObj[this.defaultProps.value])
      if (index >= 0) this.choosedValue.splice(index, 1)

      this.$refs.tree.setCheckedKeys([...this.choosedValue])
      // this.$emit('handleChange', this.choosedValue)
    },
    // 递归找到符合的元素
    findItemByName(items, targetName) {
      const {children, label} = this.defaultProps
      for (let i = 0; i < items.length; i++) {
        const currentItem = items[i]
        if (currentItem[label] === targetName) {
          return currentItem
        }

        if (currentItem[children]) {
          const foundItem = this.findItemByName(currentItem[children], targetName)
          if (foundItem) {
            return foundItem
          }
        }
      }
    },
    // 清空所有标签
    clearHandle() {
      this.allChecked = false
      this.valueName = []
      this.choosedValue = []
      this.$refs.tree.setCheckedKeys([])
      // this.$emit('handleChange', [])
    },
  },
}
</script>
<style lang="scss" scoped>
.depart-all /deep/ {
  .el-select__tags {
    max-height: 90px;
    overflow-y: auto;
  }
}
.dropdown-tree-wrap /deep/ {
  .el-checkbox__input.is-checked .el-checkbox__inner,
  .el-checkbox__input.is-indeterminate .el-checkbox__inner {
    background-color: $color !important;
    border-color: $color !important;
  }
  .el-checkbox__inner:hover,
  .el-checkbox__inner:focus,
  .el-checkbox__input.is-focus .el-checkbox__inner {
    border-color: $color !important;
  }
  .el-checkbox__input.is-checked + .el-checkbox__label {
    color: $color !important;
  }
}
</style>
