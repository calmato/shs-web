<template>
  <the-form-group>
    <v-row>
      <v-col cols="12" sm="6">
        <the-text-field
          :label="form.options.lastName.label"
          :rules="form.options.lastName.rules"
          :value.sync="form.params.lastName"
          :autofocus="true"
        />
      </v-col>
      <v-col cols="12" sm="6">
        <the-text-field
          :label="form.options.firstName.label"
          :rules="form.options.firstName.rules"
          :value.sync="form.params.firstName"
        />
      </v-col>
      <v-col cols="12" sm="6">
        <the-text-field
          :label="form.options.lastNameKana.label"
          :rules="form.options.lastNameKana.rules"
          :value.sync="form.params.lastNameKana"
        />
      </v-col>
      <v-col cols="12" sm="6">
        <the-text-field
          :label="form.options.firstNameKana.label"
          :rules="form.options.firstNameKana.rules"
          :value.sync="form.params.firstNameKana"
        />
      </v-col>
      <v-col cols="12">
        <the-text-field
          :label="form.options.mail.label"
          :rules="form.options.mail.rules"
          :value.sync="form.params.mail"
          type="email"
        />
      </v-col>
      <v-col cols="12">
        <the-text-field
          :label="form.options.password.label"
          :rules="form.options.password.rules"
          :value.sync="form.params.password"
          type="password"
        />
      </v-col>
      <v-col cols="12">
        <the-text-field
          :label="form.options.passwordConfirmation.label"
          :rules="form.options.passwordConfirmation.rules"
          :value.sync="form.params.passwordConfirmation"
          type="password"
        />
      </v-col>
      <v-col cols="12" sm="6">
        <the-select
          :label="form.options.schoolType.label"
          :rules="form.options.schoolType.rules"
          :value.sync="form.params.schoolType"
          :items="schoolTypeItem"
        />
      </v-col>
      <v-col cols="12" sm="6">
        <the-select
          :label="form.options.grade.label"
          :rules="form.options.grade.rules"
          :value.sync="form.params.grade"
          :items="getGrade(form.params.schoolType)"
        />
      </v-col>
    </v-row>
  </the-form-group>
</template>

<script lang="ts">
import { defineComponent, PropType } from '@nuxtjs/composition-api'
import TheFormGroup from '~/components/atoms/TheFormGroup.vue'
import TheSelect from '~/components/atoms/TheSelect.vue'
import TheTextField from '~/components/atoms/TheTextField.vue'
import { StudentNewForm, StudentNewOptions, StudentNewParams } from '~/types/form'
import { gradeItem, SchoolTypeItem } from '~/types/props/student'

export default defineComponent({
  components: {
    TheFormGroup,
    TheSelect,
    TheTextField,
  },

  props: {
    form: {
      type: Object as PropType<StudentNewForm>,
      default: () => ({
        params: StudentNewParams,
        options: StudentNewOptions,
      }),
    },
  },

  setup() {
    const schoolTypeItem: SchoolTypeItem[] = [
      { text: '小学校', value: '小学校' },
      { text: '中学校', value: '中学校' },
      { text: '高校', value: '高校' },
    ]

    const getGrade = (schoolType: string): gradeItem[] => {
      switch (schoolType) {
        case '小学校':
          return [
            { text: '1', value: '1' },
            { text: '2', value: '2' },
            { text: '3', value: '3' },
            { text: '4', value: '4' },
            { text: '5', value: '5' },
            { text: '6', value: '6' },
          ]
        case '中学校':
          return [
            { text: '1', value: '1' },
            { text: '2', value: '2' },
            { text: '3', value: '3' },
          ]
        case '高校':
          return [
            { text: '1', value: '1' },
            { text: '2', value: '2' },
            { text: '3', value: '3' },
          ]
        default:
          return [{ text: '校種を選択してください', value: '' }]
      }
    }

    return {
      schoolTypeItem,
      getGrade,
    }
  },
})
</script>
