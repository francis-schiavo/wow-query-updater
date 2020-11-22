package updater

import (
	blizzard_api "github.com/francis-schiavo/blizzard-api-go"
	"wow-query-updater/datasets"
)

func UpdateMount(data *blizzard_api.ApiResponse) {
	var mount datasets.Mount
	data.Parse(&mount)

	if mount.Faction != nil {
		mount.FactionID = mount.Faction.ID
	}
	if mount.Source != nil {
		mount.SourceID = mount.Source.ID
		insertOnce(mount.Source)
	}

	insertOnceUpdate(&mount, "name", "description", "faction_id", "source_id")

	if mount.MountDisplays != nil{
		for _, media := range mount.MountDisplays {
			media.MountID = mount.ID
			insertOnceUpdate(media, "mount_id")
		}
	}
}

func UpdateMountDisplayMedia(data *blizzard_api.ApiResponse, id int) {
	var mountDisplayMedia datasets.MountDisplayMedia
	data.Parse(&mountDisplayMedia)

	for _, asset := range mountDisplayMedia.Assets {
		asset.MountDisplayMediaID = mountDisplayMedia.ID
		insertOnceExpr(&asset, "(mount_display_media_id,key) DO UPDATE", "value")
	}
}