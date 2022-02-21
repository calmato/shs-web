<template>
  <v-app class="root">
    <the-snackbar :snackbar.sync="snackbar" :color="snackbarColor" :message="snackbarMessage" />
    <the-header :overlay="overlay" @click="handleClickMenu" />
    <the-sidebar :items="getMenuItems()" :current="current" @click="handleClickMenuItem" />
    <v-main>
      <nuxt />
      <the-menu
        :overlay="overlay"
        :absolute="false"
        :items="getMenuItems()"
        @click:item="handleClickMenuItem"
        @click:close="handleClickMenu"
      />
    </v-main>
  </v-app>
</template>

<script lang="ts">
import { computed, defineComponent, ref, useRoute, useRouter, useStore, watch } from '@nuxtjs/composition-api'
import TheHeader from '~/components/organisms/TheHeader.vue'
import TheMenu from '~/components/organisms/TheMenu.vue'
import TheSidebar from '~/components/organisms/TheSidebar.vue'
import TheSnackbar from '~/components/organisms/TheSnackbar.vue'
import { CommonStore } from '~/store'
import { Menu } from '~/types/props/menu'
import { Role } from '~/types/store'

export default defineComponent({
  components: {
    TheHeader,
    TheMenu,
    TheSnackbar,
    TheSidebar,
  },

  setup(_) {
    const route = useRoute()
    const router = useRouter()
    const store = useStore()

    const items: Menu[] = [
      {
        name: 'トップ',
        icon: 'mdi-home',
        path: '/',
        filter: 'all',
      },
      {
        name: 'シフト希望',
        icon: 'mdi-note-edit-outline',
        path: '/submissions',
        filter: 'all',
      },
      {
        name: 'シフト作成',
        icon: 'mdi-chart-box-plus-outline',
        path: '/shifts',
        filter: [Role.ADMINISTRATOR],
      },
      {
        name: 'ユーザ一覧',
        icon: 'mdi-format-list-bulleted',
        path: '/users',
        filter: 'all',
      },
      {
        name: '設定',
        icon: 'mdi-cogs',
        path: '/settings',
        filter: 'all',
      },
    ]

    const current = ref<string>(route.value.path)
    const snackbar = ref<Boolean>(false)
    const overlay = ref<boolean>(false)

    const role = computed<Role>(() => store.getters['auth/getRole'])
    const snackbarColor = computed(() => store.getters['common/getSnackbarColor'])
    const snackbarMessage = computed(() => store.getters['common/getSnackbarMessage'])

    watch(
      () => route,
      (): void => {
        current.value = route.value.path
      }
    )

    watch(snackbarMessage, (): void => {
      snackbar.value = snackbarMessage.value !== ''
    })

    watch(snackbar, (): void => {
      if (!snackbar.value) {
        CommonStore.hiddenSnackbar()
      }
    })

    const getMenuItems = (): Menu[] => {
      const menu = items.filter((item: Menu) => {
        if (item.filter === 'all') {
          return true
        }

        return item.filter.includes(role.value)
      })
      return menu
    }

    const handleClickMenu = (): void => {
      overlay.value = !overlay.value
    }

    const handleClickMenuItem = (item: Menu): void => {
      router.push(item.path)
      overlay.value = false
    }

    return {
      overlay,
      current,
      snackbar,
      snackbarColor,
      snackbarMessage,
      getMenuItems,
      handleClickMenu,
      handleClickMenuItem,
    }
  },
})
</script>

<style lang="css" scoped>
.root {
  overflow: hidden;
}
</style>
