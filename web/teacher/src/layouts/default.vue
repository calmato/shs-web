<template>
  <v-app class="root" :style="{ background }">
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
import { defineComponent, ref, SetupContext, watch } from '@nuxtjs/composition-api'
import { VuetifyThemeItem } from 'vuetify/types/services/theme'
import TheHeader from '~/components/organisms/TheHeader.vue'
import TheMenu from '~/components/organisms/TheMenu.vue'
import { Menu } from '~/types/props/menu'

export default defineComponent({
  components: {
    TheHeader,
    TheMenu,
  },

  setup(_, { root }: SetupContext) {
    const router = root.$router
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

    const overlay = ref<boolean>(false)
    const background = ref<VuetifyThemeItem>(getBackgroundColor(root.$route.path))

    watch(
      () => root.$route,
      (): void => {
        background.value = getBackgroundColor(root.$route.path)
      }
    )

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
