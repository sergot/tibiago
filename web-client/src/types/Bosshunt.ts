import type Participant from "./Participant"

export default interface Bosshunt {
    id: string
    boss: string
    world: string
    discord_ref: string
    discord_only: boolean
    when: string
    participants: Participant[]
}