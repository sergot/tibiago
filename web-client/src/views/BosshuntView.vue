<script setup lang="ts">
import { useRoute } from 'vue-router';
import ParticipantHolder from '../components/ParticipantHolder.vue'

import axios from 'axios'
import { computed, onMounted, ref } from 'vue';
import type Bosshunt from '@/types/Bosshunt';
import { tibianTime } from '@/utils';

const route = useRoute();
const bosshunt_id = computed(() => route.params.id);

const bosshunt = ref<Bosshunt>();

onMounted(async () => {
  try {
    const response = await axios.get('http://localhost:1323/bosshunt/' + bosshunt_id.value)
    bosshunt.value = response.data
  } catch (error) {
    console.error(error)
  }
})
</script>
<template>
  <section class="section">
    <div class="container">
      <div class="columns">
        <div class="column" v-if="bosshunt && bosshunt.participants">
          <p>{{ bosshunt.boss }} {{ tibianTime(bosshunt.when) }}</p>
          <p>List of participants</p>
          <p v-for="(participant, index) in bosshunt.participants" :key="index">
            <ParticipantHolder :vocation="participant.vocation"/>
          </p>
          test
        </div>
      </div>
    </div>
  </section>
</template>

<style></style>
