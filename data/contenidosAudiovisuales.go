package data

import (
	"errors"

	"github.com/NeichS/graphql-pixdex/graph/model"
)

// Todas las IDs, Generos y TipoID han sido convertidos a string.
var ContenidosAudiovisuales = []*model.ContenidoAudioVisual{
	{
		ID:          "1",
		Nombre:      "Breaking Bad",
		Descripcion: ptr("Un profesor de química se convierte en fabricante de metanfetaminas."),
		Generos:     []string{"1", "3", "10"},
		TipoID:      "1",
		ImageURL:    ptr("https://placehold.co/400x600"),
	},
	{
		ID:          "2",
		Nombre:      "Stranger Things",
		Descripcion: ptr("Un grupo de niños se enfrenta a fuerzas sobrenaturales en su ciudad."),
		Generos:     []string{"1", "4", "8"},
		TipoID:      "1",
		ImageURL:    ptr("https://placehold.co/400x600"),
	},
	{
		ID:          "3",
		Nombre:      "Friends",
		Descripcion: ptr("Seis amigos viven las altas y bajas de la vida en Nueva York."),
		Generos:     []string{"2", "1"},
		TipoID:      "1",
		ImageURL:    ptr("https://placehold.co/400x600"),
	},
	{
		ID:          "4",
		Nombre:      "The Office",
		Descripcion: ptr("La vida cotidiana en una oficina de papel con mucho humor absurdo."),
		Generos:     []string{"2"},
		TipoID:      "1",
		ImageURL:    ptr("https://placehold.co/400x600"),
	},
	{
		ID:          "5",
		Nombre:      "Game of Thrones",
		Descripcion: ptr("Familias nobles luchan por el control de los Siete Reinos."),
		Generos:     []string{"1", "6", "5"},
		TipoID:      "1",
		ImageURL:    ptr("https://placehold.co/400x600"),
	},
	{
		ID:          "6",
		Nombre:      "The Mandalorian",
		Descripcion: ptr("Un cazarrecompensas viaja por la galaxia protegiendo a un misterioso niño."),
		Generos:     []string{"4", "5", "11"},
		TipoID:      "1",
		ImageURL:    ptr("https://placehold.co/400x600"),
	},
	{
		ID:          "7",
		Nombre:      "Dark",
		Descripcion: ptr("Desapariciones en un pueblo alemán revelan oscuros secretos temporales."),
		Generos:     []string{"8", "1", "4"},
		TipoID:      "1",
		ImageURL:    ptr("https://placehold.co/400x600"),
	},
	{
		ID:          "8",
		Nombre:      "The Crown",
		Descripcion: ptr("Dramatización de la vida de la reina Isabel II."),
		Generos:     []string{"1", "15"},
		TipoID:      "1",
		ImageURL:    ptr("https://placehold.co/400x600"),
	},
	{
		ID:          "9",
		Nombre:      "Sherlock",
		Descripcion: ptr("Una moderna adaptación de Sherlock Holmes."),
		Generos:     []string{"8", "1"},
		TipoID:      "1",
		ImageURL:    ptr("https://placehold.co/400x600"),
	},
	{
		ID:          "10",
		Nombre:      "The Boys",
		Descripcion: ptr("Un grupo busca exponer los crímenes de superhéroes corruptos."),
		Generos:     []string{"3", "11", "10"},
		TipoID:      "1",
		ImageURL:    ptr("https://placehold.co/400x600"),
	},
	{
		ID:          "11",
		Nombre:      "Better Call Saul",
		Descripcion: ptr("El origen del abogado de Breaking Bad."),
		Generos:     []string{"1", "10"},
		TipoID:      "1",
		ImageURL:    ptr("https://placehold.co/400x600"),
	},
	{
		ID:          "12",
		Nombre:      "How I Met Your Mother",
		Descripcion: ptr("Un hombre cuenta a sus hijos cómo conoció a su madre."),
		Generos:     []string{"2", "9"},
		TipoID:      "1",
		ImageURL:    ptr("https://placehold.co/400x600"),
	},
	{
		ID:          "13",
		Nombre:      "The Witcher",
		Descripcion: ptr("Un cazador de monstruos lucha por encontrar su lugar en un mundo brutal."),
		Generos:     []string{"5", "6", "1"},
		TipoID:      "1",
		ImageURL:    ptr("https://placehold.co/400x600"),
	},
	{
		ID:          "14",
		Nombre:      "House of Cards",
		Descripcion: ptr("Un político ambicioso trama su camino hacia el poder."),
		Generos:     []string{"1", "10"},
		TipoID:      "1",
		ImageURL:    ptr("https://placehold.co/400x600"),
	},
	{
		ID:          "15",
		Nombre:      "Peaky Blinders",
		Descripcion: ptr("Una familia criminal británica después de la Primera Guerra Mundial."),
		Generos:     []string{"1", "3", "15"},
		TipoID:      "1",
		ImageURL:    ptr("https://placehold.co/400x600"),
	},
	{
		ID:          "16",
		Nombre:      "Black Mirror",
		Descripcion: ptr("Antología que explora el lado oscuro de la tecnología."),
		Generos:     []string{"4", "10"},
		TipoID:      "1",
		ImageURL:    ptr("https://placehold.co/400x600"),
	},
	{
		ID:          "17",
		Nombre:      "Euphoria",
		Descripcion: ptr("Adolescentes enfrentan amor, adicciones y traumas."),
		Generos:     []string{"1", "9"},
		TipoID:      "1",
		ImageURL:    ptr("https://placehold.co/400x600"),
	},
	{
		ID:          "18",
		Nombre:      "Succession",
		Descripcion: ptr("Una familia millonaria lucha por el control del imperio mediático."),
		Generos:     []string{"1", "10"},
		TipoID:      "1",
		ImageURL:    ptr("https://placehold.co/400x600"),
	},
	{
		ID:          "19",
		Nombre:      "Lost",
		Descripcion: ptr("Supervivientes de un accidente aéreo intentan sobrevivir en una isla misteriosa."),
		Generos:     []string{"5", "8"},
		TipoID:      "1",
		ImageURL:    ptr("https://placehold.co/400x600"),
	},
	{
		ID:          "20",
		Nombre:      "The Last of Us",
		Descripcion: ptr("Una pandemia convierte a la humanidad en monstruos."),
		Generos:     []string{"3", "4", "1"},
		TipoID:      "1",
		ImageURL:    ptr("https://placehold.co/400x600"),
	},
	{
		ID:          "21",
		Nombre:      "Inception",
		Descripcion: ptr("Un ladrón entra en los sueños de otros para robar secretos."),
		Generos:     []string{"4", "10", "1"},
		TipoID:      "2",
		ImageURL:    ptr("https://placehold.co/400x600"),
	},
	{
		ID:          "22",
		Nombre:      "The Matrix",
		Descripcion: ptr("Un hacker descubre la verdad sobre su realidad simulada."),
		Generos:     []string{"4", "3"},
		TipoID:      "2",
		ImageURL:    ptr("https://placehold.co/400x600"),
	},
	{
		ID:          "23",
		Nombre:      "Pulp Fiction",
		Descripcion: ptr("Historias entrelazadas de crimen en Los Ángeles."),
		Generos:     []string{"1", "2", "10"},
		TipoID:      "2",
		ImageURL:    ptr("https://placehold.co/400x600"),
	},
	{
		ID:          "24",
		Nombre:      "Titanic",
		Descripcion: ptr("Un amor trágico a bordo del famoso barco."),
		Generos:     []string{"1", "9", "15"},
		TipoID:      "2",
		ImageURL:    ptr("https://placehold.co/400x600"),
	},
	{
		ID:          "25",
		Nombre:      "The Dark Knight",
		Descripcion: ptr("Batman enfrenta al temible Joker en Ciudad Gótica."),
		Generos:     []string{"3", "11", "10"},
		TipoID:      "2",
		ImageURL:    ptr("https://placehold.co/400x600"),
	},
	{
		ID:          "26",
		Nombre:      "Forrest Gump",
		Descripcion: ptr("Un hombre simple vive momentos clave de la historia de EE.UU."),
		Generos:     []string{"1", "2", "15"},
		TipoID:      "2",
		ImageURL:    ptr("https://placehold.co/400x600"),
	},
	{
		ID:          "27",
		Nombre:      "Interstellar",
		Descripcion: ptr("Astronautas viajan a través de un agujero de gusano para salvar a la humanidad."),
		Generos:     []string{"4", "1", "5"},
		TipoID:      "2",
		ImageURL:    ptr("https://placehold.co/400x600"),
	},
	{
		ID:          "28",
		Nombre:      "Fight Club",
		Descripcion: ptr("Un hombre crea un club de lucha secreto que cambia su vida."),
		Generos:     []string{"1", "10"},
		TipoID:      "2",
		ImageURL:    ptr("https://placehold.co/400x600"),
	},
	{
		ID:          "29",
		Nombre:      "Gladiator",
		Descripcion: ptr("Un general romano busca venganza como gladiador."),
		Generos:     []string{"3", "15", "1"},
		TipoID:      "2",
		ImageURL:    ptr("https://placehold.co/400x600"),
	},
	{
		ID:          "30",
		Nombre:      "The Godfather",
		Descripcion: ptr("La historia de una poderosa familia mafiosa."),
		Generos:     []string{"1", "10", "15"},
		TipoID:      "2",
		ImageURL:    ptr("https://placehold.co/400x600"),
	},
	{
		ID:          "31",
		Nombre:      "Avengers: Endgame",
		Descripcion: ptr("Los Vengadores enfrentan la batalla final contra Thanos."),
		Generos:     []string{"3", "11", "4"},
		TipoID:      "2",
    ImageURL:    ptr("https://placehold.co/400x600"),
  },
}

// ptr is a helper to get a pointer to a string literal.
func ptr(s string) *string {
	return &s
}

func GetContenidoPorID(id string) (*model.ContenidoAudioVisual,error) {
	for _, contenido := range ContenidosAudiovisuales {
		if contenido.ID == id {
			return contenido, nil
		}
	}
	return nil, errors.New("contenido not found")
}

func MapContenidoAudioVisual(contenido *model.ContenidoAudioVisual) *model.ContenidoAudioVisualMapped {
	// Mapear generos
	var generosMapped []*model.Genero
	for _, generoID := range contenido.Generos {
		genero := GetGeneroPorId(generoID)
		generosMapped = append(generosMapped, genero)
	}

	// Mapear tipo
	tipo := GetTipoPorId(contenido.TipoID)

	return &model.ContenidoAudioVisualMapped{
		ID:          contenido.ID,
		Nombre:      contenido.Nombre,
		Descripcion: contenido.Descripcion,
		Generos:     generosMapped,
		Tipo:        &tipo,
		ImageURL:    contenido.ImageURL,
	}
}