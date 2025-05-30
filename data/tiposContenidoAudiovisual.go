package data

import "github.com/NeichS/graphql-pixdex/graph/model"


var tiposContenidoAudiovisual = []model.TipoContenidoAudioVisual{
  { ID: "1", Singular: "serie", Plural: "series" },
  { ID: "2", Singular: "película", Plural: "películas" },
  { ID: "3", Singular: "anime", Plural: "animes" },
}

/**
 * Busca un tipo por su id y lo retorna. En caso de no encontrarlo, retorna un tipo con el id consultado y nombre "-"
 * @param id number
 * @returns ITipoContenidoAudiovisual
 */
func GetTipoPorId(id string) model.TipoContenidoAudioVisual {
  for _, tipo := range tiposContenidoAudiovisual {
    if tipo.ID == id {
      return tipo
    }
  }
  return model.TipoContenidoAudioVisual{ ID: id, Singular: "-", Plural: "-" }
}