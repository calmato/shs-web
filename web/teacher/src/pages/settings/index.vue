<template>
  <setting-top :user-items="userItems" :system-items="systemItems" @click="handleClick" />
</template>

<script lang="ts">
import { defineComponent, SetupContext } from '@nuxtjs/composition-api'
import SettingTop from '~/components/templates/SettingTop.vue'
import { Menu } from '~/types/props/setting'

export default defineComponent({
  components: {
    SettingTop,
  },

  setup(_, { root }: SetupContext) {
    const router = root.$router

    const userItems: Menu[] = [
      {
        title: 'ユーザー情報の変更',
        path: '#ユーザー情報の変更',
      },
      {
        title: 'パスワードの変更',
        path: '#パスワードの変更',
      },
      {
        title: 'サインアウト',
        path: '#サインアウト',
      },
    ]
    const systemItems: Menu[] = [
      {
        title: '教室・科目設定',
        path: '#教室・科目設定',
      },
      {
        title: 'コマ設定',
        path: '#コマ設定',
      },
    ]

    const handleClick = (item: Menu): void => {
      router.push(item.path)
    }

    return {
      userItems,
      systemItems,
      handleClick,
    }
  },
})
</script>
