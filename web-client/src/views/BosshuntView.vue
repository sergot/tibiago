<script setup lang="ts">
import { useRoute } from 'vue-router';
import ParticipantHolder from '../components/ParticipantHolder.vue'

import axios from 'axios'
import { computed, onMounted } from 'vue';

const route = useRoute();
const bosshunt_id = computed(() => route.params.id);

onMounted(async () => {
  try {
    const response = await axios.get('http://localhost:1323/bosshunt/' + bosshunt_id.value)
    console.log(response.data)
  } catch (error) {
    console.error(error)
  }
})
</script>
<template>
  <section class="section">
    <div class="container">
      <div class="columns">
        <div class="column">
          <p>Ferumbras 12:00 12-12-12</p>
          <p>List of participants</p>
          <ParticipantHolder/>
          <ParticipantHolder vocation="ek"/>
          <!-- TODO: participant component -->
          test
        </div>
      </div>
    </div>
  </section>
</template>

<style></style>
