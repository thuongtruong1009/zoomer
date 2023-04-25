import { axiosWsInstance } from './axios'

export const StreamSevices = {
    createStream() {
        return axiosWsInstance.post('/streams/create')
    },

    joinStream(streamId: string) {
        return axiosWsInstance.post(`/streams/join/${streamId}`)
    },
}
