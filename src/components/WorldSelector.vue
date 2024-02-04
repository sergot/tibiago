<script setup lang="ts">
import { ref } from 'vue';

const world = ref('');
</script>

<template>
  <section class="section">
    <div class="container">
      <div class="columns">
        <div class="column">
          <div class="field">
            <label class="label">Name</label>
            <div class="control">
              <div class="select">
                <select v-model="world">
                <option disabled value="">Choose your world...</option>
                <option>Jaguna</option>
                <option>Kendria</option>
                <!-- more worlds from mmo game called Tibia -->
              </select>
              </div>
            </div>
          </div>
          <div class="field is-grouped">
            <div class="control">
              <RouterLink to="/world"><button class="button is-link">Select</button></RouterLink>
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
