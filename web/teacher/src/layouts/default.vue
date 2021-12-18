<template>
  <v-app class="root" :style="{ background }">
    <the-snackbar :snackbar.sync="snackbar" :color="snackbarColor" :message="snackbarMessage" />
    <the-header :overlay="overlay" @click="handleClickMenu" />
    <v-main>
      <nuxt />
      <the-menu
        :overlay="overlay"
        :absolute="true"
        :items="items"
        @click:item="handleClickMenuItem"
        @click:close="handleClickMenu"
      />
    </v-main>
  </v-app>
</template>

<script lang="ts">
import { computed, defineComponent, ref, SetupContext, watch } from '@nuxtjs/composition-api'
import { VuetifyThemeItem } from 'vuetify/types/services/theme'
import TheHeader from '~/components/organisms/TheHeader.vue'
import TheMenu from '~/components/organisms/TheMenu.vue'
import TheSnackbar from '~/components/organisms/TheSnackbar.vue'
import { CommonStore } from '~/store'
import { Menu } from '~/types/props/menu'

export default defineComponent({
  components: {
    TheHeader,
    TheMenu,
    TheSnackbar,
  },

  setup(_, { root }: SetupContext) {
    const router = root.$router
    const store = root.$store
    const vuetify = root.$vuetify

    const greyBackgroundPaths = ['/settings']
    const items: Menu[] = [
      {
        name: 'トップ',
        icon: 'mdi-home',
        path: '/',
      },
      {
        name: 'シフト希望',
        icon: 'mdi-note-edit-outline',
        path: '#シフト希望', // TODO: update
      },
      {
        name: 'シフト作成',
        icon: 'mdi-chart-box-plus-outline',
        path: '/shifts',
      },
      {
        name: 'ユーザ一覧',
        icon: 'mdi-format-list-bulleted',
        path: '/users',
      },
      {
        name: '設定',
        icon: 'mdi-cogs',
        path: '/settings',
      },
    ]

    const getBackgroundColor = (path: string): VuetifyThemeItem => {
      const theme = vuetify.theme.dark ? 'dark' : 'light'

      let color: VuetifyThemeItem = vuetify.theme.themes[theme].white
      if (greyBackgroundPaths.includes(path)) {
        color = vuetify.theme.themes[theme].grey
      }

      return color
    }

    const snackbar = ref<Boolean>(false)
    const overlay = ref<boolean>(false)
    const background = ref<VuetifyThemeItem>(getBackgroundColor(root.$route.path))

    const snackbarColor = computed(() => store.getters['common/getSnackbarColor'])
    const snackbarMessage = computed(() => store.getters['common/getSnackbarMessage'])

    watch(
      () => root.$route,
      (): void => {
        background.value = getBackgroundColor(root.$route.path)
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

    const handleClickMenu = (): void => {
      overlay.value = !overlay.value
    }

    const handleClickMenuItem = (item: Menu): void => {
      router.push(item.path)
      handleClickMenu()
    }

    return {
      items,
      overlay,
      background,
      snackbar,
      snackbarColor,
      snackbarMessage,
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
