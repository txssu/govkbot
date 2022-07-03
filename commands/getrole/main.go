package getrole

import (
	"context"
	"fmt"
	"vkbot/core"
	"vkbot/core/rolesystem"

	"github.com/SevereCloud/vksdk/v2/events"
)

func Register() core.Command {
	return core.Command{
		Aliases:     []string{"роль"},
		Description: "узнать свою роль в беседе",
		Handler:     handle,
	}
}

func handle(ctx *context.Context, obj *events.MessageNewObject) {
	s := core.GetStorage()

	initialized, _ := s.Db.Get(s.Ctx, fmt.Sprintf("roles.%d.initialized", obj.Message.PeerID)).Result()
	if initialized != "true" {
		core.ReplySimple(obj, core.ERR_NO_ROLESYSTEM)

		return
	}

	role := ""

	switch rolesystem.GetRole(obj) {
	case rolesystem.ROLE_OWNER:
		role = "владелец"
	case rolesystem.ROLE_ADMINISTRATOR:
		role = "администратор"
	case rolesystem.ROLE_MODERATOR:
		role = "модератор"
	case rolesystem.ROLE_MEMBER:
		role = "участник"
	}

	core.ReplySimple(obj, "ваша роль: "+role)
}