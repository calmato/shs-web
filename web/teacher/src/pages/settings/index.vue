<template>
  <the-setting-top :user-items="userItems" :system-items="systemItems" @click="handleClick" />
</template>

<script lang="ts">
import { defineComponent, SetupContext } from '@nuxtjs/composition-api'
import TheSettingTop from '~/components/templates/TheSettingTop.vue'
import { AuthStore } from '~/store'
import { Menu } from '~/types/props/setting'

export default defineComponent({
  components: {
    TheSettingTop,
  },

  setup(_, { root }: SetupContext) {
    const router = root.$router

    const userItems: Menu[] = [
      {
        title: 'ユーザー情報の変更',
        path: '#ユーザー情報の変更',
      },
      {
        title: 'メールアドレスの変更',
        path: '#メールアドレスの変更',
      },
      {
        title: 'パスワードの変更',
        path: '#パスワードの変更',
      },
      {
        title: 'サインアウト',
        path: '/signout',
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
      if (item.path === '/signout') {
        AuthStore.signOut()
        router.push('/signin')
        return
      }

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
