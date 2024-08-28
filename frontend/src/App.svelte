<script>
  import { GetRandomImageUrl, GetBreedList, GetImageUrlsByBreed } from '../wailsjs/go/main/App.js'

  let randomImageUrl = ''
  let breedList = []
  let imageUrlsByBreed = []
  let selectedBreed = ''
  let showImageUrlsByBreed = false
  let showRandomImage = false

  function init() {
    fetchBreedList()
  }
  init()


  async function fetchRandomImageUrl() {
    showRandomImage = true
    showImageUrlsByBreed = false
    randomImageUrl = await GetRandomImageUrl()
  } 

  async function fetchBreedList() {
    breedList = await GetBreedList()
  }

  async function fetchImageUrlsByBreed(breed) {
    init()
    showImageUrlsByBreed = true
    showRandomImage = false
    imageUrlsByBreed = await GetImageUrlsByBreed(breed)
  }
  
  
</script>

<h3>Dogs API</h3>
<div>
  <button on:click={fetchRandomImageUrl}>Random</button>
  Click on down arrow to select a breed
  <select bind:value={selectedBreed}>
    {#each breedList as breed}
      <option value={breed}>{breed}</option>
    {/each}
  </select>
  <button on:click={() => fetchImageUrlsByBreed(selectedBreed)}>Fetch By This Breed</button>
</div>
<br />
{#if showRandomImage}
  <div>    
    <img id="random-image" src={randomImageUrl} alt="Random Dog Image" aria-hidden="true" />
  </div>
{/if}

{#if showImageUrlsByBreed}
  <div>    
    {#each imageUrlsByBreed as imageUrl}
      <img id="image-by-breed" src={imageUrl} alt="Dog Image" aria-hidden="true"/>
    {/each}
  </div>
{/if}

<style>
  #random-image {
    width: 600px;
    height: auto;
  }
  #image-by-breed {
    width: 300px;
    height: auto;
  }

  button:focus {
    border-width: 3px;
  }

</style>
