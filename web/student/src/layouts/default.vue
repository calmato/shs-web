<template>
  <v-app class="root" :style="{ background }">
    <the-snackbar :snackbar.sync="snackbar" :color="snackbarColor" :message="snackbarMessage" />
    <the-header :overlay="overlay" @click="handleClickMenu" />
    <the-sidebar :items="items" :current="current" @click="handleClickMenuItem" />
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
import TheSidebar from '~/components/organisms/TheSidebar.vue'
import TheSnackbar from '~/components/organisms/TheSnackbar.vue'
import { CommonStore } from '~/store'
import { Menu } from '~/types/props/menu'

export default defineComponent({
  components: {
    TheHeader,
    TheMenu,
    TheSnackbar,
    TheSidebar,
  },

  setup(_, { root }: SetupContext) {
    const route = root.$route
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
        name: '授業希望',
        icon: 'mdi-note-edit-outline',
        path: '/submissions',
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

    const current = ref<string>(route.path)
    const snackbar = ref<Boolean>(false)
    const overlay = ref<boolean>(false)
    const background = ref<VuetifyThemeItem>(getBackgroundColor(root.$route.path))

    const snackbarColor = computed(() => store.getters['common/getSnackbarColor'])
    const snackbarMessage = computed(() => store.getters['common/getSnackbarMessage'])

    watch(
      () => root.$route,
      (): void => {
        current.value = root.$route.path
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
      overlay.value = false
    }

    return {
      items,
      overlay,
      current,
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
