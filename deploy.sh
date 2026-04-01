#!/bin/bash

# ======================================================
# 湘遇恋爱博客 - 快速部署脚本
# 使用：./deploy.sh
# ======================================================

set -e

echo "💕 湘遇恋爱博客 - 快速部署"
echo "================================"

# 颜色定义
RED='\033[0;31m'
GREEN='\033[0;32m'
YELLOW='\033[1;33m'
NC='\033[0m' # No Color

# 检查 Docker
check_docker() {
    if ! command -v docker &> /dev/null; then
        echo -e "${RED}错误：Docker 未安装${NC}"
        echo "请先安装 Docker: https://docs.docker.com/get-docker/"
        exit 1
    fi

    if ! command -v docker-compose &> /dev/null; then
        echo -e "${RED}错误：Docker Compose 未安装${NC}"
        echo "请先安装 Docker Compose"
        exit 1
    fi

    echo -e "${GREEN}✓ Docker 环境检查通过${NC}"
}

# 检查环境变量
check_env() {
    if [ ! -f .env ]; then
        echo -e "${YELLOW}! 未找到 .env 文件，从 .env.example 创建${NC}"
        cp .env.example .env
        echo -e "${YELLOW}! 请编辑 .env 文件配置数据库和 JWT_SECRET${NC}"
        echo ""
        read -p "按回车继续..."
    fi
}

# 构建镜像
build() {
    echo ""
    echo -e "${GREEN}开始构建 Docker 镜像...${NC}"
    docker-compose build --no-cache
    echo -e "${GREEN}✓ 构建完成${NC}"
}

# 启动服务
start() {
    echo ""
    echo -e "${GREEN}启动服务...${NC}"
    docker-compose up -d

    echo ""
    echo -e "${GREEN}等待服务启动...${NC}"
    sleep 10

    # 检查服务状态
    docker-compose ps
}

# 显示访问信息
show_info() {
    echo ""
    echo "================================"
    echo -e "${GREEN}🎉 部署完成！${NC}"
    echo ""
    echo "访问地址：http://localhost:5700"
    echo ""
    echo "默认管理员账户:"
    echo "  用户名：admin"
    echo "  密码：myhx0308"
    echo ""
    echo "常用命令:"
    echo "  查看日志：docker-compose logs -f"
    echo "  停止服务：docker-compose down"
    echo "  重启服务：docker-compose restart"
    echo "================================"
}

# 主流程
main() {
    check_docker
    check_env
    build
    start
    show_info
}

# 运行
main
