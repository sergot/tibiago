<script setup lang="ts">
import { ref } from 'vue';

defineProps({
  vocation: {
    type: String,
    default: 'any'
  },
})

const startedParticipating = ref(false)
const showSignUp = ref(false)
const showModal = ref(false)

const startParticipating = () => {
    showModal.value = true
}

const confirmParticipation = () => {
    startedParticipating.value = true
    showModal.value = false

    // TODO: send request to server
    // TODO: assign vocation based on user's profile instead of using any
}

const cancelParticipation = () => {
    showModal.value = false
    if (startedParticipating.value) {
        startedParticipating.value = false
        // TODO: send request to server
    }
}

const toggleSignUp = () => {
    showSignUp.value = !showSignUp.value
}
</script>

<template>
    <div v-if="!startedParticipating" class="highlight" @mouseover="toggleSignUp" @mouseout="toggleSignUp">
        <p @click="startParticipating">{{ vocation }} <span v-show="showSignUp">Click to sign up</span></p>
    </div>
    <div v-else>
        <p class="cancel" @click="cancelParticipation"><b>user</b> as {{ vocation }}</p>
    </div>

    <b-modal v-model="showModal"
    :can-cancel="['outside', 'escape']"
    @cancel="cancelParticipation"
  >
    <div class="box" style="width: 50%; margin: auto;">
      <p>Do you really want to sign up as a participant?</p>
      <button class="button is-success" @click="confirmParticipation">Yes</button>
      <button class="button is-danger" @click="cancelParticipation">No</button>
    </div>
  </b-modal>
</template>

<style scoped>
.highlight:hover {
    background-color: yellow;
    cursor: pointer;
}

.cancel:hover {
    background-color: red;
    cursor: pointer;
}
</style>