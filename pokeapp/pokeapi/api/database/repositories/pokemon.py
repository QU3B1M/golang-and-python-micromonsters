from .base import BaseRepository
from ..schemas import Pokemon


# TODO Update the BaseRepo Models as they are created
class PokemonRepository(BaseRepository[Pokemon, Pokemon, Pokemon]):
    """Repository that handles the Pokemon CRUD"""

    model: Pokemon = Pokemon
