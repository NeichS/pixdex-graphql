package data

import (
	"github.com/NeichS/graphql-pixdex/graph/model"
)

var generosContenidoAudiovisual =  []model.Genero{
  { ID: "1", Nombre: "drama" },
  { ID: "2", Nombre: "comedia" },
  { ID: "3", Nombre: "acción" },
  { ID: "4", Nombre: "ciencia ficción" },
  { ID: "5", Nombre: "aventura" },
  { ID: "6", Nombre: "fantasía" },
  { ID: "7", Nombre: "terror" },
  { ID: "8", Nombre: "misterio" },
  { ID: "9", Nombre: "romance" },
  { ID: "10", Nombre: "thriller" },
  { ID: "11", Nombre: "superhéroes" },
  { ID: "12", Nombre: "slice of life" },
  { ID: "13", Nombre: "shonen" },
  { ID: "14", Nombre: "mecha" },
  { ID: "15", Nombre: "histórico" },
  { ID: "16", Nombre: "documental" },
};

/**
 * Busca un genero por su id y lo retorna. En caso de no encontrarlo, retorna un genero con el id consultado y Nombre "-"
 * @param id number
 * @returns IGeneroContenidoAudiovisual
 */
func GetGeneroPorId(id string) *model.Genero {
  for _, genero := range generosContenidoAudiovisual{
    if genero.ID == id {
      return &genero
    }
  }
  return nil 
}