const execSync = require('child_process').execSync
const fs=require('fs');

const cmd1 = 'git init';
var cmd2 = 'git remote add -f origin ';
const cmd3 = 'git config core.sparsecheckout true';
const cmd4='git pull origin master';

var Path='';
process.stdin.setEncoding('utf8');
process.stdout.write('请输入仓库远程地址，以及想要复制的文件路径？两个输入以及多个路径都用空格分开:\n');
process.stdin.on('data',(data)=>{
	cmd2+=data.split(' ')[0];
	Path=data.slice(data.indexOf(' ')+1).split(' ').join('\n');
	process.stdin.emit('end');
});

process.stdin.on('end',()=>{
	execSync(cmd1);
	execSync(cmd2);
	execSync(cmd3);

	fs.writeFileSync('./.git/info/sparse-checkout',Path);
	execSync(cmd4)
	console.log('It\'s ok!');

});
