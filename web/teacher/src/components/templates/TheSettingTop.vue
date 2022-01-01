<template>
  <v-container class="px-0">
    <the-subject-select-form-item
      :user="user"
      :elementary-school-subjects="elementarySchoolSubjects"
      :junior-high-school-subjects="juniorHighSchoolSubjects"
      :high-school-subjects="highSchoolSubjects"
    />
    <v-row class="py-4">
      <v-col cols="12">
        <div class="text-subtitle-1">ユーザー設定</div>
        <v-card v-for="item in userItems" :key="`user-${item.title}`" elevation="0" class="my-1" @click="onClick(item)">
          <v-list-item>
            <v-list-item-content>
              <v-list-item-title>{{ item.title }}</v-list-item-title>
            </v-list-item-content>
            <v-list-item-action>
              <v-icon>mdi-chevron-right</v-icon>
            </v-list-item-action>
          </v-list-item>
        </v-card>
      </v-col>
      <v-col>
        <div class="text-subtitle-1">教室設定</div>
        <v-card
          v-for="item in systemItems"
          :key="`sys-${item.title}`"
          elevation="0"
          class="my-1"
          @click="onClick(item)"
        >
          <v-list-item>
            <v-list-item-content>
              <v-list-item-title>{{ item.title }}</v-list-item-title>
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
import TheSubjectSelectFormItem from '~/components/molecules/TheSubjectSelectFormItem.vue'
import { Menu } from '~/types/props/setting'
import { Subject } from '~/types/store'

interface UserPoop {
  lastName: string
  firstName: string
  lastNameKana: string
  firstNameKana: string
}

export default defineComponent({
  components: {
    TheSubjectSelectFormItem,
  },
  props: {
    userItems: {
      type: Array as PropType<Menu[]>,
      default: () => [],
    },
    systemItems: {
      type: Array as PropType<Menu[]>,
      default: () => [],
    },
    user: {
      type: Object as PropType<UserPoop>,
      default: () => ({
        lastName: '',
        firstName: '',
        lastNameKana: '',
        firstNameKana: '',
      }),
    },

    elementarySchoolSubjects: {
      type: Array as PropType<Subject[]>,
      default: () => [],
    },
    juniorHighSchoolSubjects: {
      type: Array as PropType<Subject[]>,
      default: () => [],
    },
    highSchoolSubjects: {
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
