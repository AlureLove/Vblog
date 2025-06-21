<template>
  <a-layout class="layout">
    <a-layout-header class="header">
      <div>Top Navigator</div>
    </a-layout-header>
    <a-layout class="body">
      <a-layout-sider width="260">
        <a-menu @menu-item-click="handleMenuItemClick">
          <a-sub-menu key="blog">
            <template #title>
              <span>
                <icon-bookmark/>Article Management
              </span>
            </template>
            <a-menu-item key="backend_blog_list">Articles</a-menu-item>
            <a-menu-item key="backend_tag_list">Tags</a-menu-item>
          </a-sub-menu>
          <a-sub-menu key="comment">
            <template #title>
              <span>
                <icon-message/>Comment Management
              </span>
            </template>
            <a-menu-item key="backend_comment_list">Comment List</a-menu-item>
          </a-sub-menu>
        </a-menu>
      </a-layout-sider>
      <a-layout-content>
        <RouterView></RouterView>
      </a-layout-content>
    </a-layout>
  </a-layout>
</template>

<script setup>

import {useRouter} from "vue-router";

const router = useRouter();

const handleMenuItemClick = (key) => {
  router.push({name: key})
}

</script>

<style lang="css" scoped>

.layout {
  height: 100%;
}

.header {
  height: 60px;
  box-sizing: border-box;
  min-height: 60px;
  max-height: 60px;
  display: flex;
  justify-content: space-between;
  position: fixed;
  top: 0;
  left: 0;
  right: 0;
  z-index: 999;
  border-bottom: 1px solid var(--color-border);
  background-color: var(--color-bg-2);
}

.body {
  margin-top: 60px;
}

</style>
