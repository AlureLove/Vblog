<template>
  <a-layout class="layout">
    <a-layout-header class="header">
      <div>Top Navigator</div>
    </a-layout-header>
    <a-layout class="body">
      <a-layout-sider width="260">
        <a-menu>
          <a-sub-menu key="1">
            <template #title>
              <span>
                <IconCalendar/>Navigation1
              </span>
            </template>
            <a-menu-item key="1_1">Menu 1</a-menu-item>
            <a-menu-item key="1_2">Menu 2</a-menu-item>
          </a-sub-menu>
          <a-sub-menu key="2">
            <template #title>
              <span>
                <IconCalendar/>Navigation2
              </span>
            </template>
            <a-menu-item key="2_1">Menu 1</a-menu-item>
            <a-menu-item key="2_2">Menu 2</a-menu-item>
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
