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
      <v-col cols="12">
        <the-select
          :label="form.options.role.label"
          :rules="form.options.role.rules"
          :value.sync="form.params.role"
          :items="roleItems"
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
import { TeacherNewForm, TeacherNewOptions, TeacherNewParams } from '~/types/form'
import { RoleItem } from '~/types/props/teacher'
import { Role } from '~/types/store'

export default defineComponent({
  components: {
    TheFormGroup,
    TheSelect,
    TheTextField,
  },

  props: {
    form: {
      type: Object as PropType<TeacherNewForm>,
      default: () => ({
        params: TeacherNewParams,
        options: TeacherNewOptions,
      }),
    },
  },

  setup() {
    const roleItems: RoleItem[] = [
      { text: '--- 役職を選択してください -------', value: Role.UNKNOWN },
      { text: '講師', value: Role.TEACHER },
      { text: '管理者', value: Role.ADMINISTRATOR },
    ]

    return {
      roleItems,
    }
  },
})
</script>
