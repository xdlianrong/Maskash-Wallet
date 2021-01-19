<template>
    <div class="button">{{ buttonMsg }}</div>
</template>
<style>
    .button {
        font-size: 1.2rem;
        text-align: center;
        background: #59aafa;
        border-radius: 4px;
        border: #f2f2f2 10px solid;
        color: white;
        line-height: 3.5rem;
    }
</style>
<script>
export default {
    props: {
        buttonMsg: {
            type: String,
            require: true
        }
    }
}
</script>