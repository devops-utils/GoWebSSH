<template>
  <div
    style="position: absolute; width: 100%; height: 100%; text-align: center"
  >
    <el-result icon="error" title="404" subTitle="兄弟,你访问的页面不存在">
      <template #extra>
        <el-link type="primary" href="/">主页</el-link>
      </template>
    </el-result>
  </div>
</template>
