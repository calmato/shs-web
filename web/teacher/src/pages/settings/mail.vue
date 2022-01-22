<template>
  <div>
    <p class="text-h5 mt-3 mb-1 ml-3">メールアドレス変更</p>
    <the-form-group>
      <v-row>
        <v-col cols="10" class="ml-3">
          <the-text-field
            :label="form.options.mail.label"
            :rules="form.options.mail.rules"
            :value.sync="form.params.updateMail"
            type="email"
          />
        </v-col>
      </v-row>
    </the-form-group>
    <v-container>
      <v-row class="justify-end">
        <v-btn color="primary" class="right mt-3 mr-15"> 保存する </v-btn>
      </v-row>
    </v-container>
  </div>
</template>

<script lang="ts">
import { defineComponent, PropType } from '@vue/composition-api'
import TheFormGroup from '~/components/atoms/TheFormGroup.vue'
import TheTextField from '~/components/atoms/TheTextField.vue'
import { TeacherUpdateMailForm, TeacherUpdateMailOptions, TeacherUpdateMailParams } from '~/types/form'

export default defineComponent({
  components: {
    TheFormGroup,
    TheTextField,
  },

  props: {
    form: {
      type: Object as PropType<TeacherUpdateMailForm>,
      default: () => ({
        params: TeacherUpdateMailParams,
        options: TeacherUpdateMailOptions,
      }),
    },
  },

  setup() {},
})
</script>
