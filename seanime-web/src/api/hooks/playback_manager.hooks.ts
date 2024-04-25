import { useServerMutation } from "@/api/client/requests"
import { PlaybackStartPlaylist_Variables } from "@/api/generated/endpoint.types"
import { API_ENDPOINTS } from "@/api/generated/endpoints"
import { useServerStatus } from "@/app/(main)/_hooks/server-status.hooks"
import { useQueryClient } from "@tanstack/react-query"
import { toast } from "sonner"

export function usePlaybackSyncCurrentProgress() {
    const serverStatus = useServerStatus()
    const queryClient = useQueryClient()

    return useServerMutation<number>({
        endpoint: API_ENDPOINTS.PLAYBACK_MANAGER.PlaybackSyncCurrentProgress.endpoint,
        method: API_ENDPOINTS.PLAYBACK_MANAGER.PlaybackSyncCurrentProgress.methods[0],
        mutationKey: [API_ENDPOINTS.PLAYBACK_MANAGER.PlaybackSyncCurrentProgress.key],
        onSuccess: async mediaId => {
            if (!serverStatus?.isOffline) {
                await queryClient.invalidateQueries({ queryKey: [API_ENDPOINTS.ANIME_ENTRIES.GetAnimeEntry.key, String(mediaId)] })
                await queryClient.invalidateQueries({ queryKey: [API_ENDPOINTS.ANIME_COLLECTION.GetLibraryCollection.key] })
                await queryClient.invalidateQueries({ queryKey: [API_ENDPOINTS.ANILIST.GetAnilistCollection.key] })
            } else {
                await queryClient.invalidateQueries({ queryKey: [API_ENDPOINTS.OFFLINE.GetOfflineSnapshot.key] })
            }
        },
    })
}

export function usePlaybackPlayNextEpisode(...keys: any) {
    return useServerMutation<boolean>({
        endpoint: API_ENDPOINTS.PLAYBACK_MANAGER.PlaybackPlayNextEpisode.endpoint,
        method: API_ENDPOINTS.PLAYBACK_MANAGER.PlaybackPlayNextEpisode.methods[0],
        mutationKey: [API_ENDPOINTS.PLAYBACK_MANAGER.PlaybackPlayNextEpisode.key, ...keys],
        onSuccess: async () => {

        },
    })
}

export function usePlaybackStartPlaylist() {
    const queryClient = useQueryClient()

    return useServerMutation<boolean, PlaybackStartPlaylist_Variables>({
        endpoint: API_ENDPOINTS.PLAYBACK_MANAGER.PlaybackStartPlaylist.endpoint,
        method: API_ENDPOINTS.PLAYBACK_MANAGER.PlaybackStartPlaylist.methods[0],
        mutationKey: [API_ENDPOINTS.PLAYBACK_MANAGER.PlaybackStartPlaylist.key],
        onSuccess: async () => {
            await queryClient.invalidateQueries({ queryKey: [API_ENDPOINTS.PLAYLIST.GetPlaylists.key] })
        },
    })
}

export function usePlaybackCancelCurrentPlaylist(...keys: any) {
    return useServerMutation<boolean>({
        endpoint: API_ENDPOINTS.PLAYBACK_MANAGER.PlaybackCancelCurrentPlaylist.endpoint,
        method: API_ENDPOINTS.PLAYBACK_MANAGER.PlaybackCancelCurrentPlaylist.methods[0],
        mutationKey: [API_ENDPOINTS.PLAYBACK_MANAGER.PlaybackCancelCurrentPlaylist.key, ...keys],
        onSuccess: async () => {
            toast.info("Cancelling playlist")
        },
    })
}

export function usePlaybackPlaylistNext(...keys: any) {
    return useServerMutation<boolean>({
        endpoint: API_ENDPOINTS.PLAYBACK_MANAGER.PlaybackPlaylistNext.endpoint,
        method: API_ENDPOINTS.PLAYBACK_MANAGER.PlaybackPlaylistNext.methods[0],
        mutationKey: [API_ENDPOINTS.PLAYBACK_MANAGER.PlaybackPlaylistNext.key, ...keys],
        onSuccess: async () => {
            toast.info("Loading next file")
        },
    })
}

