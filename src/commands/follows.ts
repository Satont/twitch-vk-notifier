import { Twitch } from '../libs/twitch'
import { config } from '../helpers/config'
import { User } from '../models/User'

const twitch = new Twitch(config.twitch.clientId)

export default async ({ userId, service}: { userId: number, service: 'vk' | 'telegram'}): Promise<number | string[]> => {
  const user = await User.findOne({ where: { id: userId, service } })
  if (!user || !user.follows.length) {
    return 0
  } else {
    const channels = await twitch.getChannelsById(user.follows)
    return channels.map(o => o.displayName)
  }
}