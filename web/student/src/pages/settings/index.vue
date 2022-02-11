<template>
  <the-setting-top class="px-2 mt-4" :menu-items="menuItems" :user="auth" @click="handleClick" />
</template>

<script lang="ts">
import { computed, defineComponent, useRouter, useStore } from '@nuxtjs/composition-api'
import TheSettingTop from '~/components/templates/TheSettingTop.vue'
import { AuthStore } from '~/store'
import { Menu } from '~/types/props/setting'
import { Auth } from '~/types/store'

export default defineComponent({
  components: {
    TheSettingTop,
  },

  setup() {
    const router = useRouter()
    const store = useStore()

    const menuItems: Menu[] = [
      {
        title: 'メールアドレス変更',
        path: '/settings/mail',
      },
      {
        title: 'パスワードの変更',
        path: '/settings/password',
      },
      {
        title: '授業希望のカスタム設定',
        path: '/settings/custom',
      },
      {
        title: 'サインアウト',
        path: '/signout',
        textColor: 'error',
      },
    ]

    const auth = computed<Auth>(() => store.getters['auth/getAuth'])

    const handleClick = (item: Menu): void => {
      console.log(item.path)
      if (item.path === '/signout') {
        AuthStore.signOut()
        router.push('/signin')
        return
      }
      router.push(item.path)
    }

    return {
      menuItems,
      handleClick,
      auth,
    }
  },
})
</script>
