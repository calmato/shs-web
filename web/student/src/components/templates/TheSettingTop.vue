<template>
  <v-container>
    <the-setting-user-form-item :user="user" :subjects="subjects" />
    <v-row class="py-4">
      <v-col cols="12">
        <v-card
          v-for="item in menuItems"
          :key="`menu-${item.title}`"
          elevation="0"
          class="my-1"
          outlined
          @click="onClick(item)"
        >
          <v-list-item>
            <v-list-item-content>
              <v-list-item-title :class="`${item.textColor}--text`">{{ item.title }}</v-list-item-title>
            </v-list-item-content>
            <v-list-item-action>
              <v-icon>mdi-chevron-right</v-icon>
            </v-list-item-action>
          </v-list-item>
        </v-card>
      </v-col>
    </v-row>
  </v-container>
</template>

<script lang="ts">
import { defineComponent, PropType, SetupContext } from '@nuxtjs/composition-api'
import TheSettingUserFormItem from '~/components/molecules/TheSettingUserFormItem.vue'
import { Menu, UserProps } from '~/types/props/setting'
import { Subject } from '~/types/store'

export default defineComponent({
  components: {
    TheSettingUserFormItem,
  },

  props: {
    menuItems: {
      type: Array as PropType<Menu[]>,
      default: () => [],
    },
    user: {
      type: Object as PropType<UserProps>,
      default: () => ({
        lastName: '',
        firstName: '',
        lastNameKana: '',
        firstNameKana: '',
      }),
    },
    subjects: {
      type: Array as PropType<Subject[]>,
      default: () => [],
    },
  },

  setup(_, { emit }: SetupContext) {
    const onClick = (item: Menu): void => {
      emit('click', item)
    }

    return {
      onClick,
    }
  },
})
</script>
