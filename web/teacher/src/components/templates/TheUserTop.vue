<template>
  <v-container class="px-0 pt-0">
    <v-row>
      <v-col cols="12">
        <v-tabs v-model="selector" grow class="pb-4">
          <v-tab v-for="actor in actors" :key="actor.value" :href="`#tab-${actor.value}`">
            {{ actor.name }}
          </v-tab>
        </v-tabs>

        <v-tabs-items v-model="selector">
          <v-tab-item value="tab-teachers">
            <the-teacher-list
              :items="teachers"
              :total="teachersTotal"
              :loading="loading"
              :page="teachersPage"
              :items-per-page="teachersItemsPerPage"
              @update:page="$emit('update:teachers-page', $event)"
              @update:items-per-page="$emit('update:teachers-items-per-page', $event)"
            />
          </v-tab-item>
          <v-tab-item value="tab-students">
            <the-student-list :items="students" :loading="loading" />
          </v-tab-item>
        </v-tabs-items>
      </v-col>
    </v-row>
  </v-container>
</template>

<script lang="ts">
import { defineComponent, PropType, ref } from '@nuxtjs/composition-api'
import TheStudentList from '~/components/organisms/TheStudentList.vue'
import TheTeacherList from '~/components/organisms/TheTeacherList.vue'
import { Actor } from '~/types/props/user'
import { Student, Teacher } from '~/types/store'

export default defineComponent({
  components: {
    TheStudentList,
    TheTeacherList,
  },

  props: {
    loading: {
      type: Boolean,
      default: false,
    },
    students: {
      type: Array as PropType<Student[]>,
      default: () => [],
    },
    teachers: {
      type: Array as PropType<Teacher[]>,
      default: () => [],
    },
    teachersTotal: {
      type: Number,
      default: 0,
    },
    teachersPage: {
      type: Number,
      default: 1,
    },
    teachersItemsPerPage: {
      type: Number,
      default: 10,
    },
  },

  setup() {
    const actors: Actor[] = [
      { name: '講師', value: 'teachers' },
      { name: '生徒', value: 'students' },
    ]

    const selector = ref<string>('teachers')

    return {
      actors,
      selector,
    }
  },
})
</script>
