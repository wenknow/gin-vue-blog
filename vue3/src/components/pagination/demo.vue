<template>
    <pagination v-model:current="current1" :total="50" show-less-items />     
    <pagination
      show-size-changer
      v-model:current="current2"
      v-model:pageSize="pageSize"
      :total="500"
      @showSizeChange="onShowSizeChange"
    />
    <br />
    <pagination
      v-model:current="current3"
      show-size-changer
      :total="500"
      disabled
      @showSizeChange="onShowSizeChange"
    />
    <pagination
    v-model:current="current4"
    :page-size-options="pageSizeOptions"
    :total="total"
    show-size-changer
    :page-size="pageSize"
    @showSizeChange="onShowSizeChange"
  >
    <template #buildOptionText="props">
      <span v-if="props.value !== '50'">{{ props.value }}条/页</span>
      <span v-else>全部</span>
    </template>
  </pagination>
  <div>
    <pagination show-quick-jumper v-model:current="current8" :total="500" @change="onChange" />
  </div>
  <div>
    <pagination
      :total="85"
      :show-total="total => `Total ${total} items`"
      :page-size="20"
      v-model:current="current6"
    />
    <br />
    <pagination
      :total="85"
      :show-total="(total, range) => `${range[0]}-${range[1]} of ${total} items`"
      :page-size="20"
      v-model:current="current7"
    />
  </div>
</template>
<script>
import pagination from "./main";
export default {
  components: {
    pagination,
  },
  data() {
    return {
      total: 10,
      pageSize: 20,
      current1: 3,
      current2: 4,
      current3: 4,
      current4: 4,
      current5: 4,
      current6: 4,
      current7: 4,
      current8: 4,
      pageSizeOptions: ['10', '20', '30', '40', '50'],
    };
  },
  methods: {
    onShowSizeChange(current, pageSize) {
      console.log(current, pageSize);
    },
    onChange(pageNumber) {
      console.log('Page: ', pageNumber);
    },
  },
};
</script>