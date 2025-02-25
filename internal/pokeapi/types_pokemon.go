package pokeapi

import "fmt"

type Pokemon struct {
    ID int
    Name string
    BaseExp int `json:"base_experience"`
    Height int
    Weight int
    IsDefault bool `json:"is_default"`
    Abilities []struct{
        IsHidden bool `json:"is_hidden"`
        Slot int
        Ability apiResource
    }
    Forms []apiResource
    HeldItems []struct{
        Item apiResource
    } `json:"held_items"`
    LocationAreaEncounters string `json:"location_area_encounters"`
    Moves []struct{
        Move apiResource
    }
    Species apiResource
    Stats []struct{
        Stat apiResource
        Effort int
        BaseStat int `json:"base_stat"`
    }
    Types []struct{
        Slot int
        Type apiResource
    }
}

func (p Pokemon) String() string {
    pstring := "pokemon{\n"
    pstring += fmt.Sprintf("ID: %d\nName: %s\nBaseExp: %d\nHeight: %d\nWeight: %d\nIsDefault: %t\n",
        p.ID, p.Name, p.BaseExp, p.Height, p.Weight, p.IsDefault,
    )
    pstring += fmt.Sprintf("Abilities: %+v\nForms: %+v\nHeldItems: %+v\nLocationAreaEncounters: %s\n",
        p.Abilities, p.Forms, p.HeldItems, p.LocationAreaEncounters,
    )
    pstring += fmt.Sprintf("Moves: %+v\nSpecies: %+v\nStats: %+v\nTypes: %+v\n",
        p.Moves, p.Species, p.Stats, p.Types,
    )
    pstring += "}"

    return pstring
}
