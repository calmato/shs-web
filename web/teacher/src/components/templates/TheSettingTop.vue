<template>
  <v-container>
    <the-subject-select-form-item
      :elementary-school-subjects-form-value.sync="elementarySchoolSubjectsFormData"
      :junior-high-school-subjects-form-value.sync="juniorHighSchoolSubjectsFormData"
      :high-school-subjects-form-value.sync="highSchoolSubjectsFormData"
      :user="user"
      :elementary-school-subjects="elementarySchoolSubjects"
      :junior-high-school-subjects="juniorHighSchoolSubjects"
      :high-school-subjects="highSchoolSubjects"
      @handleElementarySchoolSubjectsBlur="handleElementarySchoolSubjectsBlur"
      @handleJuniorHighSchoolSubjectsBlur="handleJuniorHighSchoolSubjectsBlur"
      @handleHighSchoolSubjectsBlur="handleHighSchoolSubjectsBlur"
    />
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
import { computed, defineComponent, PropType, SetupContext } from '@nuxtjs/composition-api'
import TheSubjectSelectFormItem from '~/components/molecules/TheSubjectSelectFormItem.vue'
import { Menu, UserProps } from '~/types/props/setting'
import { Subject } from '~/types/store'

export default defineComponent({
  components: {
    TheSubjectSelectFormItem,
  },

  props: {
    menuItems: {
      type: Array as PropType<Menu[]>,
      default: () => [],
    },
    systemItems: {
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

    elementarySchoolSubjectsFormValue: {
      type: Array as PropType<number[]>,
      default: () => [],
    },
    juniorHighSchoolSubjectsFormValue: {
      type: Array as PropType<number[]>,
      default: () => [],
    },
    highSchoolSubjectsFormValue: {
      type: Array as PropType<number[]>,
      default: () => [],
    },
  },

  setup(props, { emit }: SetupContext) {
    const onClick = (item: Menu): void => {
      emit('click', item)
    }

    const elementarySchoolSubjectsFormData = computed({
      get: () => props.elementarySchoolSubjectsFormValue,
      set: (val: object) => emit('update:elementarySchoolSubjectsFormValue', val),
    })

    const juniorHighSchoolSubjectsFormData = computed({
      get: () => props.juniorHighSchoolSubjectsFormValue,
      set: (val: object) => emit('update:juniorHighSchoolSubjectsFormValue', val),
    })
    const highSchoolSubjectsFormData = computed({
      get: () => props.highSchoolSubjectsFormValue,
      set: (val: object) => emit('update:highSchoolSubjectsFormValue', val),
    })

    const handleElementarySchoolSubjectsBlur = (val: number[]) => {
      emit('handleElementarySchoolSubjectsBlur', val)
    }

    const handleJuniorHighSchoolSubjectsBlur = (val: number[]) => {
      emit('handleJuniorHighSchoolSubjectsBlur', val)
    }

    const handleHighSchoolSubjectsBlur = (val: number[]) => {
      emit('handleHighSchoolSubjectsBlur', val)
    }

    return {
      onClick,
      elementarySchoolSubjectsFormData,
      juniorHighSchoolSubjectsFormData,
      highSchoolSubjectsFormData,
      handleElementarySchoolSubjectsBlur,
      handleJuniorHighSchoolSubjectsBlur,
      handleHighSchoolSubjectsBlur,
    }
  },
})
</script>
