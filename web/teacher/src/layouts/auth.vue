<template>
  <v-app class="root" :style="{ background }">
    <the-header :overlay="overlay" :show-menu="false" @click="handleClickMenu" />
    <v-main>
      <nuxt />
    </v-main>
  </v-app>
</template>

<script lang="ts">
import { defineComponent, ref, SetupContext, watch } from '@nuxtjs/composition-api'
import { VuetifyThemeItem } from 'vuetify/types/services/theme'
import TheHeader from '~/components/organisms/TheHeader.vue'
import { Menu } from '~/types/props/menu'

export default defineComponent({
  components: {
    TheHeader,
  },

  setup(_, { root }: SetupContext) {
    const router = root.$router
    const vuetify = root.$vuetify

    const getBackgroundColor = (): VuetifyThemeItem => {
      const theme = vuetify.theme.dark ? 'dark' : 'light'
      const color: VuetifyThemeItem = vuetify.theme.themes[theme].white
      return color
    }

    const overlay = ref<boolean>(false)
    const background = ref<VuetifyThemeItem>(getBackgroundColor())

    watch(
      () => root.$route,
      (): void => {
        background.value = getBackgroundColor()
      }
    )

    const handleClickMenu = (): void => {
      overlay.value = !overlay.value
    }

    const handleClickMenuItem = (item: Menu): void => {
      router.push(item.path)
      overlay.value = false
    }

    return {
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
