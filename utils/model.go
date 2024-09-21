package utils

type CliCommand struct {
	Key         string `bson:"key"`
	Name        string `bson:"name"`
	Description string `bson:"description"`
	Callback    string `bson:"callback"`
}


var	Events = map[string]CliCommand{
			"help": {
				Key:         "help",
				Name:        "You can run the MyApp using commands below:\n",
        Description: "\thelp : Lists help \n\texit : Exit from My App\n\tclear : Clear App logs\n\tshow apps : show list of Apps to use\n",
				Callback:    "commandHelp",
			},
			"show apps": {
				Key:         "show apps",
        Name:        "\nList of Apps you uan use",
    Description: "\ttodo : Use Todo for marking your Todo list \n\tword search : Use to find meanings of English words('upcoming')\n\tpassbag : Used to store ur passwords('upcoming')\n",
				Callback:    "ShowApps",
			},
			"exit": {
				Key:         "exit",
				Name:        "",
				Description: "",
				Callback:    "cmdExit",
			},
			"clear": {
				Key:         "clear",
				Name:        "",
				Description: "",
				Callback:    "cmdClear",
			},
			"todo": {
				Key:         "todo",
				Name:        "",
				Description: "",
				Callback:    "todo",
			},
			"just": {
				Key:         "just",
				Name:        "",
				Description: "",
				Callback:    "just",
			},
}
