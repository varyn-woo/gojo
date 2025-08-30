package ui

import "gojo/gen"

func MakeTextInput(id string, textField *gen.TextField) *gen.UiElement {
	return &gen.UiElement{
		Id:      id,
		Element: &gen.UiElement_TextField{TextField: textField},
	}
}

func MakeSimpleText(id string, simpleText string) *gen.UiElement {
	return &gen.UiElement{
		Id:      id,
		Element: &gen.UiElement_SimpleText{SimpleText: simpleText},
	}
}

func MakeVotingOptions(id string, votingOptions map[string]string) *gen.UiElement {
	options := []*gen.VotingOption{}
	for k, v := range votingOptions {
		options = append(options, &gen.VotingOption{
			Id:     k,
			Option: v,
		})
	}
	return &gen.UiElement{
		Id: id,
		Element: &gen.UiElement_VotingOptions{VotingOptions: &gen.VotingOptions{
			Options: options,
		}},
	}
}

func MakeStringList(id string, stringList []string) *gen.UiElement {
	return &gen.UiElement{
		Id: id,
		Element: &gen.UiElement_StringList{StringList: &gen.StringList{
			Elements: stringList,
		}},
	}
}

func MakeSimpleButton(id, label string) *gen.UiElement {
	return &gen.UiElement{
		Id:      id,
		Element: &gen.UiElement_SimpleButton{SimpleButton: label},
	}
}

func MakeCountdownTimer(id string, countdown bool) *gen.UiElement {
	return &gen.UiElement{
		Id:      id,
		Element: &gen.UiElement_CountdownTimer{CountdownTimer: countdown},
	}
}
