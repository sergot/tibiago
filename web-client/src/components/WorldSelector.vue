<script setup lang="ts">
import { computed, onMounted, ref } from 'vue';
import axios from 'axios';

import World from '@/types/World';

const worlds = ref<World[]>([]);
const world = ref('');

const selected = ref();

const filteredWorlds = computed(() => {
  return worlds.value.filter((option: World) => {
    return option.name.toLowerCase().indexOf(world.value.toLowerCase()) >= 0
  });
});

onMounted(async () => {
  try {
    const response = await axios.get('http://localhost:1323/worlds');
    worlds.value = response.data;
  } catch (error) {
    console.error('Failed to fetch worlds:', error);
  }
});
</script>

<template>
  <section class="section">
    <div class="container">
      <div class="columns">
        <div class="column">
          <b-field label="World">
            <b-autocomplete
              v-model="world"
              :data="filteredWorlds"
              placeholder="Choose your world..."
              field="name"
              :open-on-focus="true"
              @select="(option: World) => (selected = option)"
            ></b-autocomplete>
          </b-field>
          <div class="field is-grouped">
            <div class="control">
              <RouterLink :to="`/world/${world}`"><button class="button is-link">Select</button></RouterLink>
            </div>
          </div>
        </div>
        <div class="column">
          Selected world: {{ world }}
        </div>
      </div>
    </div>
  </section>
</template>
