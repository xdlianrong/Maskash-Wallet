<template>
    <div style="height: 100%;">
        <navmenu @changecmp="changecmps" ref="nav"></navmenu>  
        <el-row type="flex" justify="center" id="o">
            <el-col :xs="20" :sm="15" :md="12" :lg="8" :xl="7">  
            <div v-if="cmp == 1">
                <p>购买金额:</p>
                <el-input maxlength="10" v-model="money" ></el-input>
                <mybutton :buttonMsg="buy" @click.native="transferm"></mybutton>
            </div>

            <div v-if="cmp == 2" style="top: 10%;">
                <p>收款方公钥:</p>
                <el-input maxlength="10" v-model="money" ></el-input>
                <p>转账金额:</p>
                <el-input maxlength="10" v-model="money" ></el-input>
                <p>找零金额:</p>
                <el-input maxlength="10" v-model="money" ></el-input>
                <p>承诺:</p>
                <el-input maxlength="10" v-model="money" ></el-input>
                <p>承诺随机数:</p>
                <el-input maxlength="10" v-model="money" ></el-input>
                <mybutton :buttonMsg="transfer" @click.native="transferm"></mybutton>
            </div>

            <div v-if="cmp == 3">
                <p>shoukuan:</p>
                <el-input maxlength="10" v-model="money" ></el-input>
                <mybutton :buttonMsg="recv" @click.native="transferm"></mybutton>
            </div>

            <div v-if="cmp == 4">
                <p>转账金额:</p>
                <el-input maxlength="10" v-model="money" ></el-input>
                
            </div>
            </el-col>
        </el-row>
    </div>
</template>
<script>
import navmenu from '../components/Navmenu'
import mybutton from '../components/Mybutton'
export default {
    components: {
        navmenu,
        mybutton
    },
    data() {
        return {
            transfer: '发起转账',
            buy: '购买',
            recv: '收款',
            money: '',
            cmp: '1', // 用来改变显示的组件
        }
    },
    methods: {
        transferm() {
            console.log("我要转账");
        },
        buym() {
            console.log("我要购币");
        },
        recm() {
            console.log("我要收款");
        },
        changecmps(index) {
            this.cmp = index;
            console.log(this.cmp);
        }
    }
}
</script>
<style>

</style>