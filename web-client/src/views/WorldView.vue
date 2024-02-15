<script setup lang="ts">
import { computed, onMounted, ref } from 'vue'
import axios from 'axios'

import { useRoute } from 'vue-router'

import type Bosshunt from '@/types/Bosshunt'
import { tibianTime } from '@/utils'

const bosshunts = ref<Bosshunt[]>([]);

const route = useRoute();
const world = computed(() => route.params.world);

onMounted(async () => {
  try {
    const response = await axios.get('http://localhost:1323/bosshunts/' + world.value);
    bosshunts.value = response.data;
  } catch (error) {
    console.error('Failed to fetch bosshunts:', error);
  }
});
</script>

<template>
  <section class="section">
    <div class="container">
      <div class="columns">
        <div class="column">
          <p>List of boss hunts</p>
          <div v-for="bosshunt in bosshunts" :key="bosshunt.id">
            <RouterLink :to="`/bosshunt/${bosshunt.id}`">{{ bosshunt.boss }} [{{ tibianTime(bosshunt.when) }}] <i class="fa-brands fa-discord"></i> <i v-if="!bosshunt.discord_only" class="fa-solid fa-globe"></i></RouterLink>
          </div>
          <p>and so on...</p>
        </div>
        <div class="column">
          <p>Based on these discord servers:</p>
        </div>
      </div>
    </div>
  </section>
</template>

<style></style>
